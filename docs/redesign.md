# SILK 재설계 (상세 설계 초안, 2026-10-31)

> 기준안은 메인 제안 `files/silk/design.md`다. 이 문서는 그것을 상세화한다. 동의하지 않는 곳은 **반론**으로 표시하고 근거를 붙였다.
> 코드는 바꾸지 않았다. 근거는 세 가지다: 현 구현(`pkg/…`), 이미 잰 실측(README, `docs/compaction.md`, `files/silk/README.md`, `files/pack/conversion-2026-10-29.md`), 그리고 이번에 직접 읽은 문법 아카이브 `0x0000`과 testdata 레코드.
> 확정 설계가 아니다. 마지막 절이 사용자 결정 목록이다.

## 0. 요약: 쟁점별 권고

| # | 쟁점 | 권고 | 기준안과 다른 점 |
|---|---|---|---|
| 1 | 정렬 키·플래너 | `create index NAME TAG -part F -key F,F… -width W,W…`. 플래너는 "등호로 고정된 키 앞부분이 가장 긴 색인"을 고른다. 맞는 색인이 없으면 **거부하고** `-explain`에 만들 색인을 제안한다. `-scan`을 주면 전수 스캔한다 | 기준안은 경고 후 전수 스캔. 전체 아카이브에서는 몇 시간짜리라 기본은 거부 |
| 2 | 마스크 표기 | 별도 MASK 인자와 니블 `?`를 둘 다 받는다. 단 **비교는 필드 단위**다. 패턴은 패턴 자신의 TAG WORD로 필드를 나눈 뒤 필드마다 (값, 마스크)로 바꾼다 | 기준안의 "레코드 바이트 기준"은 같은 필드의 바이트 위치가 레코드마다 달라서 그대로는 안 된다(§5.2 실례) |
| 3 | 가변 길이 필드 | 이름 질의는 gogeul 표기(`"김치"@ko`, `Q42`, `0x…`)로 값을 쓰고 필드 조건으로 컴파일한다. `silkg encode`가 컴파일 결과를 보여 준다. 텍스트 키는 텍스트 레코드 전체가 아니라 **텍스트 바이트 하위 필드**(`object.bytes`)에 둔다 | 기준안은 경로 `2`를 키로 예시. 그 키는 길이순으로 정렬돼 접두·범위가 안 된다 |
| 4 | 범위 | `PATH=LO..HI`. 등호로 고정된 키 앞부분 바로 다음 필드 하나에 허용한다. 수·시간은 바이트 순서가 값 순서와 같은 필드만 받는다 | 기준안은 "맨 앞 필드만". 일반화해도 비용이 없다. 대신 부호 있는 날짜는 문법 보강이 필요하다(§2.4) |
| 5 | 정규화 | SILK는 바이트 그대로(기준안 동의). 응용은 먼저 **질의 쪽 변형 생성 + 일괄 질의**로 풀고, 그걸로 안 될 때 파생 GEUL 문서를 낸다. 단 문법 자체가 정한 동치(SHORT-CODE ≡ BCP 47)는 SILK 질의 컴파일러가 OR로 펼친다 | 순서와 예외를 더했다 |
| 6 | AND·결과 단위·일괄 | 문서 단위 AND(병합 조인) 유지. 결과는 `-docs`·`-records`·`-keys`(덮는 색인: pack을 읽지 않음). 일괄은 표준 입력 한 줄에 질의 하나, Go API `RunBatch` | `-keys`를 더했다. zstd에서 중요하다 |
| 7 | zstd 임의 접근 | **새 프레임 설계가 필요 없다.** gogeul pack은 이미 문서(레코드)마다 zstd 프레임 하나다. SILK는 문서 표에 인코딩을 더하고, 후보 문서를 한 번씩 풀면 된다 | 기준안의 "프레임 크기 실측"은 pack 형식 문제라 SILK 몫이 아니다 |
| 8 | 저장·명령 | 색인 선언은 매니페스트 세대 안에 둔다(`decl/<sha256>.json` + 매니페스트의 `indexes`). 파티션·런·합치기·문서 표·비트맵은 그대로 재사용 | 기준안은 별도 `catalog.json` + `idx/<NAME>/`. 세대 교체·수거와 원자적으로 묶으려면 매니페스트 안이 낫다 |
| 9 | 반증 기준 | 지금 실측을 유지하는지로 잰다. "색인 ≤ 압축 원본 50%"는 지금 이미 122%라 출발부터 기각된다. 원본 무압축 바이트 대비와 행당 바이트로 바꾼다 | 기준 수치 교체 |

## 1. 지금 구조 (코드로 확인한 사실)

| 층 | 지금 | 재설계에서 걸리는 점 |
|---|---|---|
| 명세 `pkg/spec` | `spec.json` 한 장. `Spec.Table()`이 TAG WORD(13비트) → 항목 번호 표라 **TAG당 항목 하나**만 된다 | TAG당 색인 여럿 |
| 레이아웃 | `LayoutOf`에 내장 표 7종. 필드 분류 3가지(Ref, Any, LangText) | `0x0000`에서 읽기 |
| 경로 | `[f]` 또는 `[f,0]`(언어 레코드)만 | 하위 필드 일반화, 날값 필드 |
| 키 | 필드 레코드 앞 W바이트(0 채움), W ≤ 32. 파티션 키 하나 + 정렬 키 하나 | 정렬 키를 여러 필드로 |
| 세그먼트 `pkg/seg` | 행 = 정렬 슬롯 · 문서 u32 · 오프셋 u32. `Select(part, sort)`는 접두 이분 탐색 | 범위 선택(`lo..hi`) 추가 |
| 색인 `pkg/index` | 세대 매니페스트, 문서 표(`.sdoc`, 문서마다 (TAG WORD, 레코드 수) 쌍), 비트맵, 증분, 합치기, pin | 그대로 재사용 |
| 질의 `pkg/query` | 조건 = TAG + 필드 (경로, 값 바이트). 경로가 파티션·정렬 경로와 같을 때만 키로 쓰고, 나머지는 정밀 확인. 문서 단위 병합 조인 | 플래너 재작성 |
| 아카이브 `pkg/archive` | `groupDocs`가 `Enc != raw`면 `ErrCompressed`. `Index.Record`는 pack 파일에 `ReadAt` | zstd 읽기 |

실측(16,696문서, 271 MB 무압축, 행 7,068,002개):

| 항목 | 값 |
|---|---:|
| 빌드 | 3.0–3.3 s, 113.1 MB(세그먼트 111.1 MB) |
| 행당 | 평균 15.7 B(TRIPLE 14.6, REFERENCE-NODE 24.0) |
| `P106=Q169470`(55문서) | 116–124 µs |
| `P31=Q5`(1,579문서) | 0.75–1.15 ms |
| `-label ko`(9,746건) | 3.7–4.3 ms, 대부분 pack에서 레코드를 읽어 확인하는 시간 |
| 같은 세트 zstd pack | 92.8 MB(비율 0.336) |

## 2. 문법 아카이브 `0x0000`에서 이름과 레이아웃 읽기

### 2.1 읽는 방법 (실제로 해 봄)

스크래치 프로그램으로 `draft/geul-org/artifacts/geul/.well-known/geul/0x0000/`을 읽었다.

1. `pack.Open(ctx, pack.FileStore{Root: <0x0000 세트 디렉터리>}, pack.Options{})` → `Reader.Paths`. 경로 433개.
2. Codec 규격 문서는 `0x0000/0x000000000000<TAG>/`(문서 ID = 대상 TAG WORD, 문법 `codec-spec.md`). 이번 판에 104개.
3. 최신 판을 `Reader.Fetch` → `inspect.Decode` → `codecspec.Collect(body)`. 필드 목록(이름, 폭, CODEC)이 나온다.
4. **TAG 이름**은 DECLARE Edge Id에 묶인 BIND의 UTF8-TEXT다. 문서마다 DECLARE가 하나다.

같은 일을 gogeul이 이미 한다: `resolve.Dir{Store: <pack.Archive>}.CodecSpec()`. 판 목록(`versions.geul`)의 내용 해시로 검증까지 한다. `pack.OpenDir`가 돌려주는 `*pack.Archive`가 `resolve.Store`를 구현한다. SILK는 이 경로를 쓴다.

확인한 레이아웃(일부, `필드(폭; CODEC)`, 폭 0 = 레코드, 65535 = 나머지 날값):

| TAG | 이름 | 필드 |
|---|---|---|
| `1003` | BIND | subject(0) name(0; 0840,0841) |
| `1004` | LABEL | subject(0) object(0; 084A,084B,0008) |
| `100A` | TRIPLE | predicate(0) subject(0) object(0; 0780,0F80,0001,0002) rank(1; 0802) |
| `100B` | QUALIFIER | property(0) target(0; 0008) value(0; 0780,0F80,0001,0002) |
| `1020` | SOURCE | target(0; 0008) origin(0; 0780,0001,0000) |
| `1822` | CREATED-AT (Meta) | time(0; 0820,0822,0824…0827) |
| `084A` | UTF8-LANGUAGE-TEXT | language(0; 0848,0846,0847) bytes(65535) |
| `0820` | GREGORIAN-DATE | days(2) |
| `0020` | WIKIDATA-QID | qid(65535) |

- Edge Id는 FIELD로 선언하지 않는다(모든 Edge의 암묵 첫 필드). SILK는 필드 0으로 둔다(지금과 같음).
- Meta(HEAD) TAG도 같은 방식으로 선언돼 있다. 색인 대상에 넣을 수 있다.

### 2.2 이름 ↔ ID

- 이름 → ID는 **1:N**이다. 홀짝 쌍이 같은 이름을 쓴다: UTF8-TEXT `0840/0841`, URI, EDTF, GEDCOM-DATE, LANGUAGE-BCP47-CODE, UTF8-LANGUAGE-TEXT. Edge 이름은 이번 판에서 모두 하나씩이었다.
- 그래서 `create index UTF8-TEXT`는 두 TAG WORD를 한 색인에 묶는다. 질의도 두 TAG의 OR다.
- ID 표기는 **대상 TAG WORD 4자리**(`0x1004`)만 받는다. 기준안의 `0x04`는 Edge LABEL(`1004`)인지 Reference(`0004`)인지 모호하다. 규격 문서 ID와도 맞는다.

### 2.3 문법 판 고정

- 문법 출처: ARCHIVE 안의 `0x0000` 세트가 있으면 그것(지금 위키데이터 아카이브에도 들어 있다). 없으면 `-grammar DIR`.
- 색인 선언은 그 TAG 규격 문서의 (문서, 판, 내용 해시)를 기록한다. 문법이 새 판으로 바뀌어도 색인은 기록한 판으로 계속 읽는다. `verify`가 "새 판 있음"을 알린다. 규격의 바이너리 레이아웃은 개정하지 않는 것이 문법 규칙이라(`codec-spec.md` 규칙) 판이 바뀌어도 대부분 이름·ENUM만 바뀐다.

### 2.4 `0x0000`이 모자란 곳 (발견)

| # | 모자란 것 | 근거 | 영향 | 제안 |
|---|---|---|---|---|
| G1 | TOULMIN Edge(`0x60`–`0x7B`) 규격 문서가 없다 | 문법 페이지 `edge/toulmin/`은 있으나 104개 중에 없음. 반복 구성은 DECLARE로 아직 못 쓴다(`declare.md` 열린 문제 "반복") | 이 TAG들은 필드 키를 못 만든다 | 규격 문서가 생길 때까지 TAG 단위 색인만(키 없음) |
| G2 | 날값의 부호를 모른다 | GREGORIAN-DATE `days`, POSIX-TIMESTAMP-SECONDS `seconds`는 문법 페이지상 2의 보수지만 규격 문서 FIELD에 CODEC이 없다. TRIPLE `rank`는 CODEC `0802`가 있다 | 1970년 이전 날짜는 바이트 순으로 뒤로 간다. 범위 질의가 틀린다 | 문법 보강: 해당 날값 필드에 CODEC `0801`(SIGNED-INTEGER). 이미 있는 장치("날값 필드에 CODEC 하나 = 그 Value로 해석")다 |
| G3 | codecspec이 Codec 이름을 버린다 | `codecspec.Spec`은 `map[TAG][]Field`. DECLARE의 BIND 이름은 `bindNames`(비공개)에서 쓰고 버린다 | 이름 → ID 해석을 SILK가 다시 짜야 한다 | gogeul 보강 후보: `Spec.Names()` 또는 `Collect`가 이름도 돌려주기 |
| G4 | Issuer 형식 Codec | `codecspec.Of`가 Issuer 레코드를 건너뛴다 | 응용 Issuer의 자체 Edge는 색인 못 함 | 문법 열린 문제. 이번 범위 밖 |
| G5 | 문법이 정한 동치 | SHORT-CODE `n`과 레지스트리 tag의 BCP 47(대소문자 무시)은 "같은 언어"(`language-short-code.md`) | `@ko`가 바이트 두 벌 이상 | §10.3 |

## 3. 필드 경로와 키

### 3.1 경로

- 숫자: `1`, `2.0`, `2.1`. 이름: `predicate`, `object.language`, `object.bytes`. 둘은 같은 것을 가리킨다. 이름은 FIELD의 BIND다.
- 0은 Edge Id. 1부터 선언 순서.
- 하위 경로는 레코드 필드 안으로 내려간다. 실제 레코드의 Codec 규격에 그 하위 필드가 있을 때만 값이 있다. 없으면 0 슬롯이다(지금 LABEL Object가 LOCAL일 때와 같다).
- 날값 필드도 경로가 된다. 지금은 TRIPLE `rank`를 거부하지만, 레이아웃을 알면 앞 레코드들을 TAG WORD로 건너뛰고 읽을 수 있다.
- 나머지 날값(폭 65535)은 바깥 레코드 길이에서 뒤 고정 날값을 빼서 자른다(문법 `field.md`).

### 3.2 키 슬롯 인코딩

| 필드 종류 | 슬롯 | 같음이 정확한가 |
|---|---|---|
| 레코드 필드 | 레코드 앞 W바이트, 0 채움(지금과 같음) | 레코드 ≤ W면 정확(TAG WORD에 길이가 있으므로). 넘치면 후보 |
| 고정 날값 n WORD | 2n바이트 그대로. CODEC이 부호 있는 정수면 최상위 비트를 뒤집는다(순서 보존) | 정확 |
| 나머지 날값(텍스트 바이트 등) | 앞 W−1바이트, 0 채움, 마지막 1바이트 = `00` 다 들어감 / `FF` 잘림 | `00`이면 정확, `FF`면 후보 |

- 홀수 바이트 텍스트의 끝 패딩은 문법상 `00`이다(`utf8-text.md`). 0 채움과 겹쳐서 순서가 깨지지 않는다.
- 정렬 키가 여러 필드면 슬롯을 이어 붙인다. 각 슬롯 폭이 고정이라 바이트 비교 = 필드별 사전순 비교다.
- 모든 슬롯이 "정확"이고 조건이 모두 키로 덮이면 **정밀 확인을 건너뛸 수 있다**(§6.3 덮는 색인).

### 3.3 텍스트는 `object.bytes`에 키를 둔다 (기준안 보충)

testdata LABEL 레코드의 Object 부분(실제 hex):

```
0c4a 08c6 6373 5361 6d75 …              (짧은 텍스트, 언어 cs)
0dca 08c6 6361 5361 6d75 …              (한 WORD 더 긴 텍스트)
0f4a 0015 08c6 6172 d8b5 …              (Length 확장 WORD가 끼어든다)
```

- 경로 `2`(Object 레코드 전체)를 키로 하면 맨 앞이 TAG WORD의 길이 비트라 **길이순**으로 정렬된다. 그 뒤가 언어, 그 뒤가 텍스트다.
- 그래서 정확 일치는 되지만 W 중 2–6바이트를 TAG·언어에 쓰고, 접두 일치·범위는 안 된다.
- `object.bytes`(경로 `2.1`)에 키를 두면 텍스트 바이트 순이 된다. 접두 질의(`"김치"`로 시작)가 범위 하나로 덤으로 생긴다.

## 4. 명령 체계

```sh
silkg create index NAME TAG [-part PATH] [-key PATH,PATH…] [-width W,W…] [-where PATH=VALUE…] [-min N] ARCHIVE
silkg drop index NAME ARCHIVE
silkg list index ARCHIVE                     # 이름, TAG, 키, 행 수, 바이트, 문법 판
silkg index build|update|compact|verify|stat ARCHIVE   # 지금 기능, 모든 색인 대상
silkg query [-index NAME] [-docs|-records|-keys] [-explain] [-scan] [-batch] COND… ARCHIVE
silkg encode COND…                           # 필드 조건·패턴·마스크·키 공간 범위를 출력
```

- `TAG`는 이름(`LABEL`) 또는 대상 TAG WORD(`0x1004`).
- `NAME`을 비우면(`-`) TAG와 키에서 만든다. 예: `label.object-language.object-bytes`.
- `-width`를 생략하면 필드 종류로 정한다: Reference 레코드 6, 언어 레코드 16, 텍스트 32, 고정 날값은 그 폭.
- `-where`(부분 색인): 빌드 때 조건에 맞는 레코드만 행으로 넣는다. 예: 한국어 LABEL만. §6.4.
- `-min`: 지금의 `min_partition`.
- 지금의 `-triple`·`-match`·`-label`은 공표 전이므로 지운다. 이름 질의가 같은 일을 한다(`-triple P106=Q169470` = `TRIPLE predicate=P106 object=Q169470`).

### 4.1 색인 선언 규칙

| 규칙 | 이유 |
|---|---|
| `-part`는 길이가 한정된 필드만: Reference 레코드, 언어 레코드, 고정 날값 | 파티션 표 키가 고정 폭이어야 한다(지금과 같음) |
| `-key`의 각 경로는 레이아웃으로 도달할 수 있어야 한다 | 문법에 없는 위치를 거부 |
| 같은 (TAG, part, key, width, where)는 두 번 만들지 않는다 | 선언 해시가 같다 |
| 키 없는 색인(TAG만)도 허용 | TOULMIN처럼 레이아웃 없는 TAG. "이 TAG가 든 문서" 목록 |

### 4.2 기본 색인 (지금 기본 spec을 옮김)

| 이름 | TAG | part | key | 폭 |
|---|---|---|---|---|
| triple | TRIPLE | predicate | object | 6 / 6 |
| qualifier | QUALIFIER | property | value | 6 / 6 |
| label, alias, description | LABEL, ALIAS, DESCRIPTION | object.language | subject | 16 / 6 |
| source | SOURCE | — | origin | 6 |
| reference-node | REFERENCE-NODE | — | object | 16 |

koparse가 더할 색인(예):

```sh
silkg create index label-ko LABEL -part object.language -key object.bytes,subject -width 32,6 -where object.language=@ko ARCHIVE
silkg create index alias-ko ALIAS -part object.language -key object.bytes,subject -width 32,6 -where object.language=@ko ARCHIVE
```

### 4.3 참조(Reference) TAG 색인 — 해석이 필요하다

사용자 목표는 "어떤 참조 TAG, 어떤 Edge TAG든"이다. 그런데 문서 최상위 레코드는 Meta와 Edge뿐이고, Reference는 언제나 그 **필드 안**에 있다. 그래서 Reference TAG 색인의 뜻을 정해야 한다.

| 안 | 뜻 | 행 |
|---|---|---|
| (가) 등장 색인 | 그 Codec의 Reference가 어느 레코드의 어느 깊이에든 나오면 행 하나. "Q42를 언급한 모든 레코드" | part = 감싼 최상위 TAG WORD(한정됨), key = Reference 바이트, 오프셋 = 감싼 최상위 레코드 |
| (나) 허용 안 함 | Edge 필드 경로로만 찾는다 | — |

권고: (가). 파티션이 감싼 TAG라서 "TRIPLE 안의 Q42"와 "어디든 Q42"가 같은 색인에서 나온다. 어느 필드였는지는 정밀 확인이 찾는다. 사용자 결정.

## 5. 질의

### 5.1 정규형

```
Query   = Cond+                        (같은 문서 안 AND)
Cond    = TAG 집합 + FieldPred*         (같은 레코드 안 AND)
FieldPred = (경로, 값, 마스크) | (경로, LO, HI)
값      = 한 벌 또는 여러 벌의 OR        (§10.3 언어 동치, 홀짝 Codec)
```

입력 형식은 모두 이 정규형으로 컴파일된다. `silkg encode`가 정규형과, 고른 색인의 키 공간 범위를 출력한다.

### 5.2 패턴과 마스크 — 반론: 레코드 바이트 위치는 고정이 아니다

testdata TRIPLE 두 개(실제 hex):

```
160a 0088 00b9 00a8 0241 …         predicate(00a8 0241)가 바이트 6에 있다
170a 0026 0088 00ba 00a8 0356 …    Length 확장 WORD 때문에 바이트 8에 있다
```

- Edge의 TAG WORD에 길이 비트가 있다. 14 WORD를 넘으면 Length 확장 WORD가 들어간다.
- Edge Id(LOCAL)는 1 WORD 이상 가변이다. 위키데이터 Q·P 참조도 값에 따라 1–2 WORD다.
- 그래서 "레코드 바이트 + 마스크"를 글자 그대로 적용하면 같은 필드가 레코드마다 다른 자리에 있어 맞지 않는다.

권고: 패턴은 **레코드 hex로 쓰되, 비교는 필드 단위**다.

1. 패턴 자신의 TAG WORD로 필드를 나눈다. 나눌 수 있으려면 각 필드 TAG WORD의 Kind·Codec·Length 비트는 가리면 안 된다(가리면 오류).
2. 바깥 TAG WORD는 Kind·Codec만 본다. 길이 비트와 Length 확장은 자동으로 무시한다.
3. Edge Id 필드는 마스크로 다 가려도 되고, 생략 표기도 받는다.
4. 필드마다 (값, 마스크)가 된다. 마스크가 전부 0인 필드는 조건에서 빠진다.

- 니블 와일드카드 `?`는 마스크 설탕이다. `0x…????…`와 `0xPATTERN 0xMASK`를 둘 다 받는다. 둘을 함께 주면 오류.
- 필드 길이가 다른 레코드를 한 패턴으로 잡을 수는 없다. 그건 필드 조건(이름 질의)으로 한다.
- 키 공간 패턴(색인 키 바이트로 쓰기)은 사용자 입력으로 받지 않는다. 색인을 바꾸면 질의가 달라지기 때문이다. `encode`·`-explain` 출력에만 쓴다.

### 5.3 이름 질의 (가변 길이 필드)

```sh
silkg query LABEL object.language=@ko object.bytes='"김치"' ARCHIVE
silkg query LABEL 2='"김치"@ko' ARCHIVE                  # 레코드 전체 정확 일치
silkg query TRIPLE predicate=P106 object=Q169470 ARCHIVE
silkg query TRIPLE predicate=P31 rank=1 ARCHIVE         # 날값 필드
```

- 값은 **gogeul 표기**(`geul show`·`geul add`가 쓰는 토큰)로 쓴다: `Qn`, `Pn`, `#n`, `"text"`, `"text"@tag`(BCP 47), `"text"@@tag`(SHORT-CODE), `0x…`(레코드 hex), 소수. SILK 전용 토큰을 늘리지 않는다(기준안 원칙 1과 맞음).
- 하위 필드 값(`object.bytes`)은 그 필드의 형식으로 맞춘다: 텍스트 바이트면 따옴표 안 UTF-8 바이트.
- 지금 SILK의 `encodeToken`은 gogeul 표기를 흉내 낸 복제다(`//ff:why`: gogeul `parseRecord` 비공개). gogeul이 표기 파서를 공개하면 지운다. gogeul 보강 후보.

### 5.4 범위

- `PATH=LO..HI`(양끝 포함), `PATH=LO..`, `PATH=..HI`. 값은 위 표기 또는 hex.
- 허용 자리: 등호로 고정된 키 앞부분 **바로 다음 필드 하나**. 그 뒤 키 필드는 정밀 확인으로 거른다.
  - 기준안 "맨 앞 필드만"은 이 규칙의 특수한 경우다. `TRIPLE -part predicate -key subject,object`에서 `subject=Q42 object=LO..HI`도 범위 한 번이다.
- 접두: `object.bytes^='"김치"'` = 범위 [접두, 접두+FF…]. 텍스트 키에서 덤으로 된다.
- **값 순서 = 바이트 순서인 필드만 받는다.**

| 형식 | 바이트 순 = 값 순? | 처리 |
|---|---|---|
| Reference ID(같은 Codec, 같은 길이) | 같은 WORD 수 안에서만 | 길이가 섞이면 TAG WORD 길이 비트가 먼저라 틀림. ID 범위는 거부, 등호만 |
| 고정 날값, 부호 없음(CODEC 0802 등) | 예 | 허용 |
| 고정 날값, 2의 보수(GREGORIAN-DATE, POSIX 초) | 아니오(음수가 뒤로) | 슬롯에서 부호 비트 뒤집기. 단 G2 문법 보강이 먼저 |
| 가변 정수(SIGNED/UNSIGNED-INTEGER) | 아니오(길이가 TAG에) | 거부 |
| BINARY-FLOAT, DECIMAL | 아니오 | 거부(변환은 가능하나 Codec별 코드가 필요) |
| 텍스트 바이트 | 바이트 사전순 | 허용(사전순임을 명시) |

- 위키데이터 시간 값이 실제로 어느 Codec으로 들어가는지는 확인하지 않았다(전체 변환 진행 중이라 wikigeul을 읽지 않음). 날짜 범위가 필요하면 먼저 확인한다.

### 5.5 여러 조건, 결과 단위, 일괄

- 같은 레코드 AND = 한 Cond의 필드 여러 개. 같은 문서 AND = Cond 여러 개, 문서 순번 병합 조인(지금 그대로).
- 레코드 사이 값 조인(LABEL.subject = TRIPLE.subject 같은 것)은 이번 범위 밖이다. 응용이 일괄 질의 두 번으로 한다.
- 결과 단위:

| 옵션 | 출력 | pack 읽기 |
|---|---|---|
| `-docs`(기본) | `<Issuer>/<문서>/<판>.geul #<Edge Id>,…` | 후보마다 |
| `-records` | 문서, 레코드 오프셋, Edge Id, 레코드 hex | 후보마다 |
| `-keys` | 문서, 오프셋, 키 필드 값(슬롯에서) | **없음**(덮는 색인일 때만. 아니면 오류) |

- 일괄: `-batch`이면 표준 입력 한 줄에 질의 하나. 출력은 줄 번호를 앞에 붙인다. Go API는 `query.RunBatch(ix, []Query) [][]Result`.
  - 같은 색인을 쓰는 질의는 키 순으로 정렬해 한 커서로 훑는다. 이분 탐색 probe의 페이지가 겹친다. 이득은 아직 재지 않았다.

## 6. 플래너

### 6.1 고르기

1. Cond의 TAG 집합에 걸린 색인을 모은다.
2. 파티션 필드가 있는 색인은 그 필드가 등호로 고정됐을 때만 쓴다. 파티션 없는 색인은 늘 후보.
3. 점수 = 등호로 고정된 키 앞부분 길이(필드 수). 그다음 필드에 범위가 있으면 +½.
4. 같은 점수면 행 수가 적은 색인(stat), 그다음 슬롯 폭이 좁은 색인.
5. 덮이지 않은 필드는 정밀 확인에서 거른다(지금처럼 gogeul 디코드 후 필드 비교).
6. `-where`가 있는 부분 색인은 질의가 그 조건을 포함할 때만 쓴다.
7. 값이 OR(여러 벌)이면 벌마다 선택을 돌려 합치고 (문서, 오프셋)으로 중복을 지운다.

### 6.2 맞는 색인이 없을 때 — 반론: 기본은 거부

- 파티션만 안 고정된 경우: 모든 파티션에서 키 접두를 찾는다(지금 `Select`가 하는 일). 허용.
- 그 TAG의 색인이 하나도 없으면: 문서 표의 (TAG WORD, 레코드 수) 쌍으로 그 TAG가 든 문서만 골라 읽을 수는 있다. 그래도 전체 위키데이터(zstd 약 190 GB 추정)에서는 사실상 전수 읽기다.
- 권고: 기본은 **거부**하고, `-explain`이 `silkg create index …` 명령을 제안한다. `-scan`을 주면 전수 스캔한다. 표본 아카이브에서 결과 대조(정확성 시험)에 `-scan`을 쓴다.

### 6.3 덮는 색인

- 조건 필드가 모두 키에 있고, 해당 슬롯이 모두 "정확"(§3.2)이면 정밀 확인을 건너뛴다.
- `-keys`는 행에서 바로 답한다. pack을 열지 않는다. zstd pack에서 문서 풀기를 피하는 주된 길이다.
- Edge Id가 필요하면(`-docs`) 레코드를 읽어야 한다. Edge Id를 키 끝 필드로 넣으면 그것도 피한다(`-key object.bytes,subject,0`).

### 6.4 부분 색인 `-where`

- koparse는 한국어 LABEL·ALIAS만 필요하다. 표본에서 LABEL 1,216,235행 중 `ko`는 9,746건(0.8%)이다.
- `-key object.bytes,subject -width 32,6`이면 행 = 38 + 8 = 46 B. 전체 언어면 약 56 MB, `ko`만이면 약 0.45 MB(표본 기준 계산값, 미측정).
- 파티션(언어)으로도 같은 행만 읽지만, 색인 크기와 빌드·합치기 비용은 `-where`만 줄인다.

## 7. 저장

```
ARCHIVE/.silk/
  CURRENT
  manifest/<gen>.json     # indexes: [{name, decl, runs, appends, dead}], docs, live, sets …
  decl/<sha256>.json      # 색인 선언(정규 JSON) — 지금 spec/<sha256>.json 자리
  grammar/<sha256>.json   # 선언이 쓴 TAG 규격 문서 (주소, 판, 내용 해시)와 풀어 둔 필드 목록
  seg/<n>.sseg            # 지금 형식 그대로. 꼬리의 해시 칸 = 선언 해시
  doc/<n>.sdoc, live/<n>.sbit, pin/, LOCK, compact.json   # 그대로
```

- **반론(작음):** 기준안은 `catalog.json` + `idx/<NAME>/`. 그러나 create·drop·update·compact가 모두 "새 매니페스트 세대 + CURRENT 교체" 한 번으로 원자적이어야 하고, 수거기는 세대가 가리키지 않는 파일을 지운다. 목록을 매니페스트 밖 가변 파일에 두면 그 원자성이 깨진다. 그래서 색인 목록은 매니페스트에, 선언은 내용 이름 파일에 둔다. `list index`가 사람이 볼 목록을 낸다.
- 파일 번호 카운터와 `seg/`는 하나로 둔다. 색인별 디렉터리는 이득이 없다.
- 나중에 선언을 GEUL 문서로 쓰는 일은 지금 TODO와 같다(주소 + 내용 해시로 이름).

## 8. 갱신·합치기와 재사용 범위

| 작업 | 동작 |
|---|---|
| `create index` | 살아 있는 모든 판을 한 번 읽어 그 색인만 정렬 런으로 만든다. 지금 "spec 해시가 바뀐 TAG는 모든 판에서 다시 만든다"와 같은 경로 |
| `drop index` | 매니페스트에서 뺀다. 수거기가 세그먼트를 지운다 |
| `index build` | 문서를 한 번 읽어 모든 색인의 행을 뽑는다(지금 구역별 추출과 같음) |
| `index update` | 새·옮긴 판만 읽어 색인마다 덧붙인 세그먼트 하나 |
| `index compact` | 색인마다 지금 정책(major·flush·tier) 그대로 |
| `index verify` | 색인마다 지금 검사 + 선언의 문법 판이 아카이브에 있는지 |
| `index stat` | 색인마다 행, 바이트, 덧붙인 비율, 낡은 비율 |

| 패키지 | 재사용 | 바꿀 것 |
|---|---|---|
| `pkg/seg` | 행 형식, 런, 덧붙임, k-way 병합, 체크포인트 전부 | 범위 선택 `SelectRange(part, lo, hi)` 추가 |
| `pkg/index` | 세대, 문서 표, 비트맵, pin, 수거, 합치기 정책 | "구역 = TAG" → "구역 = 색인". TAG → 색인 목록 표(지금 `Table()`은 TAG당 하나) |
| `pkg/spec` | `RecLen`, `AppendSlot`, 검증 틀 | `LayoutOf` 내장 표 → `0x0000` 규격. 경로 일반화. 다중 키, 날값 슬롯 |
| `pkg/query` | 정밀 확인, 병합 조인, 결과 | 컴파일러(이름·hex·니블·범위), 플래너, 일괄 |
| `pkg/archive` | 세트·pack·문서 목록 | zstd 허용(§9) |
| `pkg/sig` | 그대로 | 이번 범위 밖 |

- 문서 표의 (TAG WORD, 레코드 수) 쌍은 색인이 아니라 TAG 기준이라 색인을 더해도 그대로 쓴다. `dead` 계산도 그대로다(부분 색인은 행 수를 따로 세야 한다: 선언에 `-where`가 있으면 죽은 판의 해당 행을 세그먼트에서 센다).

## 9. 압축 pack

### 9.1 사실

- gogeul pack은 **레코드(= 판 문서 하나)마다 zstd 프레임 하나**다(`files/pack/README.md` R5, §zstd). 위치는 `Loc{Obj, Off, Len, Enc}`, 사전은 DICTS(위키데이터 세트 1개, 약 110 KiB).
- 그러니 "zstd 프레임 단위 임의 접근"은 이미 있다. 단위가 문서다.
- 지금 SILK는 문서 표에 (pack, Off, Len)을 갖고 있다. 행의 오프셋은 문서 안 오프셋이다. 압축 문서면 "풀린 문서 안 오프셋"으로 두면 된다. 빌드는 어차피 풀린 문서를 읽는다.

### 9.2 바꿀 것

1. `groupDocs`가 `Enc = 1`을 받는다. 문서 표 항목에 인코딩 1바이트를 더한다(예약 칸 사용 가능).
2. `Index.Record`: 압축 문서면 `[Off, Off+Len)`을 읽고 사전으로 풀어 문서 전체를 얻은 뒤 오프셋에서 레코드를 자른다.
3. 후보는 이미 (문서, 오프셋) 순으로 정렬된다. 문서마다 한 번만 푼다. 작은 LRU 캐시를 둔다.
4. gogeul 보강 후보: `Reader.record(ctx, Loc)`가 비공개다. SILK는 경로가 아니라 Loc을 안다. `Reader.Record(ctx, Loc)` 공개가 필요하다(경로 `Fetch`는 디렉터리 조회가 한 번 더 든다).

### 9.3 비용 (재지 않은 것은 가설로 적는다)

| 항목 | 근거 | 예상 |
|---|---|---|
| 문서 하나 풀기 | pack 조회 실측: 열린 Reader 전 경로 p50 6.8 µs(무압축) → 9.4 µs(zstd), p99 183 → 291 µs | 후보 문서당 수 µs–수백 µs |
| 빌드 | `check --archive` 19.6 s → 23.7 s(+21%) | 3 s → 4 s 안팎(가설) |
| `-label ko`류 대량 후보 | 지금도 읽기가 대부분(3.7 ms) | 2배를 넘을 수 있다. 덮는 색인(`-keys`)으로 피한다 |

- **반론:** 기준안 §11 "프레임 크기를 실측해서 정한다"는 pack 형식(gogeul)의 결정이다. SILK가 프레임을 따로 만들면 원본 pack 옆에 사본이 생긴다. SILK는 pack의 문서 프레임을 그대로 쓰고, 더 빨라야 하면 덮는 색인으로 읽기 자체를 없앤다.
- 무압축 사본(`GEULCOMPRESS=none`)을 계속 받을지는 사용자 결정이다. 받으면 지금 경로가 그대로 빠르다.

## 10. 정규화

### 10.1 선택지

| 안 | 내용 | 장점 | 단점 |
|---|---|---|---|
| A. 파생 GEUL 문서(기준안) | 응용이 정규화한 LABEL 등을 자기 Issuer 문서로 내고 SILK가 색인 | SILK 단순. 출처 추적(DERIVED-FROM Meta `1831`이 문법에 있음). 다른 응용도 씀 | 저장 증가. 원본이 바뀌면 다시 만들기. 위키데이터 전체 규모면 변환 한 벌 |
| B. SILK 정규화 키 | 선언에 변환 함수(`-fold space`) | 응용 코드 없음 | SILK가 유니코드 규칙을 품는다. 키가 GEUL 바이트가 아니게 된다. 정밀 확인도 정규화 비교로 바뀌고 `verify`가 복잡해진다 |
| C. 질의 쪽 변형 + 일괄 | 응용이 질의 문자열의 변형을 만들어 일괄로 묻는다 | 저장·재빌드 없음. 지금 바로 가능 | 변형 수가 폭발하면 못 쓴다 |

### 10.2 권고

- SILK는 바이트 그대로(기준안 동의, B 기각).
- 응용 순서: **C 먼저, 안 되면 A.**
  - koparse R5는 이미 "세 음절 이상 한글 어간에 공백 한 칸" 변형을 만든다(`draft/koparse/design.md`). 입력 `발효식품`(4음절)에서 라벨 `발효 식품`을 찾는 변형은 3개다. 일괄 질의 세 줄이다.
  - 반대 방향(입력에 공백, 라벨에 없음)은 공백을 지운 한 줄이다.
  - 공백 두 칸 이상이나 다른 정규화(NFC, 가운뎃점)까지 가면 변형이 늘어난다. 그때 A로 간다.
- A를 할 때는 `-where`로 그 Issuer 문서만 색인할 수 있다.

### 10.3 예외: 문법이 정한 동치

- 정규화와 달리 언어 동치는 **문법 자체의 규칙**이다: SHORT-CODE `n` ≡ 레지스트리가 `n`에 묶은 BCP 47 tag(대소문자 무시).
- wikigeul에는 SHORT-CODE 모드가 있다(geul-dev 커밋 `cecd03ee6`). 문서마다 언어 인코딩이 다를 수 있다.
- 권고: 질의 컴파일러가 `@ko`를 {SHORT-CODE `n`, BCP 47 `ko`}의 OR로 펼친다. 레지스트리는 `0x0000`의 LANGUAGE-SHORT-CODE 규격 문서 ENUM(`codecspec.Spec.Languages()`)에서 읽는다.
- 대소문자 변형(`KO`, `Ko`)은 펼치지 않는다. 소문자 정규형만 맞춘다. 대문자로 쓴 문서는 놓친다(문서화). 사용자 결정.
- 명시 표기(`"…"@ko`만, `"…"@@ko`만)를 원하면 `-exact-lang`.

## 11. 반증 기준

표본: 지금과 같은 16,696문서(271 MB) 세트, 무압축과 zstd 두 벌. 전체 아카이브는 변환이 끝난 뒤 따로 잰다.

| 가설 | 채택 | 지금 기준값 | 기각 시 |
|---|---|---|---|
| 색인을 여러 개 둬도 기본 색인 성능이 그대로다 | `P106=Q169470` ≤ 0.2 ms, `P31=Q5` ≤ 1.2 ms(지금 흔들림 0.75–1.15), 기본 7개 빌드 ≤ 5 s | 0.12 ms, 0.75–1.15 ms, 3.0–3.3 s | 색인 목록·TAG→색인 표 구조 재검토 |
| 기본 색인 크기가 늘지 않는다 | 기본 7개 ≤ 120 MB, 무압축 원본 대비 ≤ 0.5 | 113.1 MB, 0.42 | 레이아웃 해석이 키를 바꿨는지 확인 |
| 색인 하나의 크기를 예측할 수 있다 | 실측 ≤ 행 수 × (키 폭 + 8) × 1.1 | 행당 평균 15.7 B | 세그먼트 형식 재검토 |
| 텍스트 정확 일치가 빠르다 | `label-ko`로 한 단어 ≤ 1 ms, 일괄 1,000줄 ≤ 100 ms(무압축, `-keys`) | (없음) | 키 폭·일괄 커서 재설계 |
| zstd로도 쓸 만하다 | 같은 질의 지연 ≤ 무압축의 2배, 빌드 ≤ 2배, `-keys` 질의는 차이 없음 | 무압축 값 | 문서 캐시, 안 되면 무압축 사본 유지 |
| 레이아웃을 `0x0000`에서 읽어 모든 선언된 TAG를 다룬다 | 규격 문서가 있는 Edge·Meta TAG 전부에 `create index`(레코드 필드 키 1개) 성공 | — | 막히는 TAG부터 문법 보강 |
| 색인 결과가 맞다 | 무작위 질의 1,000개에서 색인 결과 = `-scan` 결과(바이트 같음), 증분·합치기 뒤에도 | 지금 4개 질의로 확인 | 버그. 공개 전 차단 |

- **반론:** 기준안의 "색인 합계 ≤ 원본(압축 기준) 50%"는 지금 이미 113.1 MB ÷ 92.8 MB = **122%**다. 출발점에서 기각되는 기준이라 판별력이 없다. 무압축 원본 대비(지금 0.42)와 행당 바이트로 바꾼다.
- **반론:** 기준안의 "문법의 모든 Edge·참조 TAG에 색인 생성 성공"은 TOULMIN(G1) 때문에 지금 이미 거짓이다. "규격 문서가 있는 TAG"로 좁혔다.

## 12. 구현 순서

| 단계 | 내용 | 끝 조건 |
|---|---|---|
| 1 | 문법 로더(`0x0000` → 이름·필드), 경로 해석. `LayoutOf` 대체 | 기본 spec 결과가 지금과 바이트 같음 |
| 2 | TAG당 색인 여럿, 매니페스트 `indexes`, create·drop·list | 증분·합치기 시험 전부 통과 |
| 3 | 다중 키, 날값·나머지 날값 슬롯, 범위 선택 | 범위·접두 결과 = `-scan` |
| 4 | 질의 컴파일러(gogeul 표기, hex 패턴, 니블), 플래너, `-explain`, `encode` | 무작위 질의 대조 |
| 5 | 일괄, 덮는 색인 `-keys`, 부분 색인 `-where` | koparse 반증 기준 |
| 6 | zstd 읽기 | zstd 반증 기준 |
| 7 | Reference 등장 색인(§4.3, 결정되면) | — |

gogeul 쪽 보강 후보(각각 작은 공개 API):

1. `codecspec`: Codec 이름(DECLARE의 BIND) 노출.
2. `pack.Reader.Record(ctx, Loc)` 공개.
3. `notation`: 레코드 토큰 하나를 읽는 함수 공개(지금 `parseRecord` 비공개).
4. 바이트 수준 레코드 프레이밍(`RecLen`, 지금 SILK `pkg/spec`에 있음, `//ff:why`에 이미 후보로 적힘).

## 13. 기준안과 다른 점 (반론 모음)

1. 패턴은 레코드 바이트로 쓰되 필드 단위로 비교한다. 레코드 바이트 위치는 Length 확장과 가변 Edge Id 때문에 고정이 아니다(§5.2 실례).
2. 텍스트 키는 `object.bytes`에. 레코드 전체(`2`)는 길이순 정렬이라 접두·범위가 안 된다(§3.3).
3. 맞는 색인이 없으면 기본 거부, `-scan`으로만 전수 스캔(§6.2).
4. 범위는 "맨 앞 필드"가 아니라 "등호 앞부분 바로 다음 필드". 부호 있는 날짜는 문법 보강(G2)이 먼저(§5.4).
5. TAG ID는 4자리 대상 TAG WORD만. `0x04`는 모호하다(§2.2).
6. zstd는 새 프레임 설계 없이 pack의 문서 프레임을 쓴다. 프레임 크기는 SILK 몫이 아니다(§9).
7. 색인 목록은 매니페스트 세대 안에. 별도 `catalog.json`은 세대 교체와 원자성이 깨진다(§7).
8. 크기 기준 "압축 원본 50%"는 지금 이미 122%라 바꾼다. "모든 TAG" 기준은 TOULMIN 때문에 좁힌다(§11).
9. 정규화는 응용 몫(동의). 다만 순서는 질의 변형 → 파생 문서. 문법이 정한 언어 동치는 SILK가 펼친다(§10).
10. 부분 색인 `-where`와 덮는 색인 `-keys`를 더한다(§6.3, §6.4).

## 14. 사용자 결정 필요

1. 패턴의 기준: "레코드 hex로 쓰고 필드 단위로 비교"(권고) / 레코드 바이트 그대로 / 색인 키 공간.
2. 마스크 표기: 별도 MASK 인자와 니블 `?`를 둘 다 받을지(권고: 둘 다, 함께 쓰면 오류).
3. 범위 허용 자리: "등호 앞부분 바로 다음 필드"(권고) / "맨 앞 필드만"(기준안).
4. 맞는 색인이 없을 때: 거부 + `-scan`(권고) / 경고 후 스캔(기준안).
5. 정규화: SILK 밖(권고·기준안). 응용 순서 질의 변형 → 파생 문서.
6. 언어 동치: `@ko`를 SHORT-CODE·BCP 47의 OR로 펼칠지(권고), 대소문자 변형은 펼치지 않을지.
7. Reference TAG 색인의 뜻: 등장 색인(권고) / 허용 안 함.
8. 문법 보강: GREGORIAN-DATE `days`·POSIX 초 등 2의 보수 날값에 CODEC `0801`을 붙일지(범위 질의의 전제).
9. TOULMIN 등 규격 문서 없는 TAG: TAG 단위 색인만 둘지, 규격 문서를 먼저 만들지.
10. 색인 목록 저장: 매니페스트 안(권고) / 별도 `catalog.json`(기준안). 선언을 처음부터 GEUL 문서로 할지(권고: JSON으로 시작).
11. 기존 `-triple`·`-match`·`-label` 플래그를 지울지(권고: 지움, 공표 전).
12. 무압축 사본을 계속 받을지(받으면 지금 성능 유지, 안 받으면 zstd 경로가 필수).
13. 기본 색인: 지금 7개를 그대로 옮길지(권고), koparse용 `label-ko`·`alias-ko`를 기본에 넣을지(권고: 넣지 않고 koparse가 만든다).
14. gogeul 보강 4건(§12)을 gogeul에 넣을지.
15. 반증 기준 수치(§11) 확정.
