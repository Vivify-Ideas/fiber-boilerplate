<h1>Fiber Boilerplate
  <a
    href="https://gofiber.io/"
    target="blank"
  >
    <img
      src="https://raw.githubusercontent.com/gofiber/docs/master/static/fiber_v2_logo.svg"
      height="25"
      alt="Fiber Logo"
    />
  </a>
</h1>

## Description

[Fiber](https://github.com/gofiber/fiber) Boilerplate made with ❤️ by [VivifyIdeas](https://www.vivifyideas.com).

## Start Guide

Just run already prepared bash script:
```bash
./init.sh
```
It will setup the project for you (starting docker-compose stack, running migrations).
The Fiber app running in dev mode will be exposed on `http://localhost` (port 3000)


## ORM integration

For more details check
[GORM](https://gorm.io/index.html)

## Environment Configuration

All environment variables stored inside `.env` file, created automatically with init bash script. Intial values copied from `.env.example`.


## Error Reporting

Integrated [sentry.io](https://sentry.io/).

Just update `SENTRY_DSN` env variable and integration is done.


## Some features

- Basic boilerplate structure
- API/Web Routes
- CORS
- Requests Validation [Validator](https://github.com/go-playground/validator)
- Database configuration [GORM](https://gorm.io/index.html)
- Authentication (login, registration, password recovery, jwt auth middleware)
- Static assets (render HTML pages)
- Notifications (sending emails)
- Errors reporting (sentry.io)
- Files upload
- Dockerized
