# drain

A log drain for Heroku's [Logplex](https://devcenter.heroku.com/articles/logplex). Parses requests and pushes them to redis.

[![Circle CI](https://circleci.com/gh/ryanlower/drain.svg?style=svg)](https://circleci.com/gh/ryanlower/drain)

---

### Usage

Deploy, then add the `/drain` endpoint as a drain to your heroku app.

For example, if drain is deployed to heroku as the app `my-new-drain` and you want to recieve logs from your `my-application`:
```
heroku drains:add https://my-new-drain.herokuapp.com/drain -a my-application
```

#### Security

To prevent unauthorized logging to your drain, you probably want to setup basic auth.

Set the `AUTH_PASSWORD` environment variable on your drain:
```
heroku set AUTH_PASSWORD=super_secret -a my-new-drain
```
And include your password in the drain url:
```
heroku drains:add https://user:super_secret@my-new-drain.herokuapp.com/drain -a my-application
```

---

### Deployment

The simplest method is to deploy on heroku:

```
git clone git@github.com:ryanlower/drain.git
cd drain
heroku apps:create -b https://github.com/kr/heroku-buildpack-go.git
git push heroku master
```

---

### Configuration

Via environment variables:

* **PORT**: The port to listen for incoming logs on (set automatically on heroku)
* **AUTH_PASSWORD**: [optional] password for basic auth for draining
* **REDIS_ADDRESS**: [optional] address of redis server, defaults to localhost:6379
* **REDIS_PASSWORD**: [optional] password for authentication to a password protected redis server

---

### Roadmap
* Add a basic UI (for 'live' view of requests)
* Add extra reporter types, e.g. send average connect time to librato
