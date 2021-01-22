# giregi.R.I.P
[archive.org](http://web.archive.org/)에 저장한 정보를 정리/관리하기 위한 어플리케이션입니다.

## Feature
- [x] archive.org 아카이브 생성
- [x] 사이트 지원
  - [x] http://news.v.daum.net (기자, 신문사, 제목, 작성일 등)
  - [x] https://news.naver.com (기자, 신문사, 제목, 작성일 등)
  - [x] https://www.clien.net (작성자, 제목, 작성일)
  - [x] http://ddanzi.com (작성자, 제목, 작성일)
- [x] 검색 (태그 검색 / keyword 검색)
- [x] 웹용
  - [x] dockerize
- [x] 데스크탑용 (feat. [electron](https://www.electronjs.org/))
  - [x] MacOS
  - [x] Windows
  - [ ] Linux(build error ㅠㅠ)
- [ ] 삭제 or 숨기기 기능
- [ ] export / backup 기능

## 빌드 환경 설정
- [go](https://golang.org/) (v1.15 이상)
- [node](https://nodejs.org/en/) (v10 이상)
- [npm](https://www.npmjs.com/)
- xcode (Mac용 빌드를 위해 필요)

## 데스크탑용 빌드 방법
[electron](https://www.electronjs.org/)으로 사이트를 데스크탑에서 실행 가능하도록 합니다.
_**빌드는 Mac에서만 빌드 가능합니다.**_
* 빌드 방법
```
npm install
npm run electron-build:mac # 맥용
npm run electron-build:win # 윈도우용
```

## 웹용 빌드 & 환경 설정
#### 웹용 빌드(Docker Build)
* server 빌드
```
cd server
docker build -t giregi.rip-server:v1.0.0 .
```
* frontend 빌드 (firebase 설정 후 [웹 앱의 구성 스니펫 가져오기](https://support.google.com/firebase/answer/7015592?hl=ko)를 참조해서 `firebaseConfig`를 json으로 저장합니다.)
```
cd frontend
cp <your firebaseConfig.json> ./src/config/google-config.json
docker build -t giregi.rip-frontend:v1.0.1 .
```

#### 서버 환경 설정
* `mysql` / `redis` / `rabbitMQ`가 설치되어 있어야 합니다.
* [config.yaml](https://github.com/ruinnel/giregi.rip/blob/master/server/config.yaml)에 mysql / redis / rabbitMQ 접속 정보를 입력 합니다.
  * `firebaseAdminJsonPath`: `firebase 콘솔 > 설정 > 서비스 계정 > Firebase Admin SDK`에서 `새 비공개키 생성`으로 생성한 json 파일.
* `docker-compose.yml` 예시
```yaml
version: "3"

services:
  mariadb:
      image: mariadb:10.5.8
      environment:
        - MYSQL_ROOT_PASSWORD=maria
        - MYSQL_DATABASE=test
      volumes:
        - ./data:/var/lib/mysql
      ports:
        - "3306:3306"
      restart: always
  redis:
    image: redis
    command: ["redis-server", "--port", "6379"]
    restart: always
  rabbitmq:
    image: rabbitmq:3-management-alpine
    volumes:
      - ./rabbitmq/etc/:/etc/rabbitmq/
      - ./rabbitmq/data/:/var/lib/rabbitmq/
      - ./rabbitmq/logs/:/var/log/rabbitmq/
    ports:
      - 5672:5672
      - 15672:15672
    environment:
      - RABBITMQ_ERLANG_COOKIE=40d958b0-2c00-4dac-93fe-9154c5f1b2ba
      - RABBITMQ_DEFAULT_USER=giregi_rip
      - RABBITMQ_DEFAULT_PASS=giregi_rip
    restart: always
  giregi-rip-server:
    image: giregi.rip-server:1.1.0
    ports:
      - "8000:8000"
    volumes:
      - ./config:/app/config
    command: ["-mode=server", "-config=/app/config/config.yaml"]
    depends_on:
      - rabbitmq
      - redis
    restart: always
  giregi-rip-worker:
    image: giregi.rip-server:1.1.0
    volumes:
      - ./config:/app/config
    command: ["-mode=worker", "-config=/app/config/config.yaml"]
    depends_on:
      - rabbitmq
      - redis
    restart: always
  frontend:
    image: giregi-rip-frontend:1.1.0
    ports:
      - "80:80"
    depends_on:
      - giregi-rip-server
    restart: always
```