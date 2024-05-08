# gc-backend-admin

## 태스크 라이브러리 사용

### 기본 태스크 라이브러리 설치 혹은 업그레이드

* tasks 디렉토리가 없을 경우 아래 명령으로 tasks를 설치 한다.
  ```bash
  task update_tasks
  ```
* 이후 해당 tasks를 커밋 하여 리포지터리에 서브모듈을 연결한다.

### 태스크 라이브러리 제거

* 템플릿 리포지터리 작성시 커밋 전에 서브모듈을 제거하기 위해 사용한다. (현재 Gitea가 지원하지 않음)
  ```bash
  task clean_tasks
  ```

### 개발 도구 설치

* 태스크 라이브러리 설치 후에 태스크 작동에 필요한 개발 도구들을 설치 할 수 있다.
  ```bash
  task install:tools
  ```
