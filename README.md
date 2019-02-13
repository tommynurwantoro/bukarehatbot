# BukaRehat Bot

Telegram bot for BukaRehat apps

## Owner

[Tommy Nurwantoro](https://github.com/tommynurwantoro)  
[Satrio Wisnugroho](https://github.com/satriowisnugroho)  
[Himang Sharatun](https://github.com/himangSharatun)  

## Onboarding and Development Guide

### Prerequisite
- Git
- Go 1.11 or later
- Go Dep 0.5 or later
- MySQL 5.7
- Redis 4.0.11 or later

### Installation

- Install Git  
  See [Git Installation](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

- Install Go (Golang)  
  See [Golang Installation](https://golang.org/doc/install)

- Install Go Dep  
  See [Dep Installation](https://golang.github.io/dep/docs/installation.html)

- Install MySQL  
  See [MySQL Installation](https://www.mysql.com/downloads/)
  
- Install Redis  
  See [Redis Installation](https://redis.io/topics/quickstart)

- Clone this repo in your local at `$GOPATH/src/github.com/bot`
  If you have not set your GOPATH, set it using [this](https://golang.org/doc/code.html#GOPATH) guide.
  If you don't have directory `src`, `github.com`, or `bot` in your GOPATH, please make them.

  ```sh
  git@github.com:tommynurwantoro/bukarehatbot.git
  ```

- Go to Kingsman directory, then sync the vendor file

  ```sh
  cd $GOPATH/src/github.com/bot/bukarehatbot
  make dep
  ```

- Copy env.sample and db/config.yml.sample if necessary, modify the env and config value(s)

  ```sh
  cp env.sample .env
  cp db/config.yml.sample db/config.yml
  ```

- Install Bundler
  
  ```sh
  gem install bundler
  ```

- Prepare database

  ```sh
  bundle install
  rake db:create db:migrate
  ```

- To create migration table, you can use this command
  ```sh
  rake db:new_migration name=foo_bar_migration
  edit db/migrate/20081220234130_foo_bar_migration.rb
  ```

- Run Bot

  ```sh
  make run
  ```
  
### Deployment
  
- Copy bukarehat.service to /lib/systemd/system/
  ```sh
  cp bukarehat.service /lib/systemd/system/bukarehat.service
  ```

- Run bukarehat
  ```sh
  service bukarehat start
  ```
  
- Stop bukarehat (If needed)
  ```sh
  service bukarehat stop
  ```
