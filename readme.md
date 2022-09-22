![is-email-disposable](https://user-images.githubusercontent.com/698079/98451232-e3ef1c80-2154-11eb-8812-0cd346fc9c6d.png)

A REST API for checking if an e-mail address is disposable (a.k.a. throwaway).

## Using the hosted version

[WebGazer](https://www.webgazer.io) offers this as a publicly available service. You can freely use it on `https://isemaildisposable.webgazer.io`. 

For example, a `GET` request to `https://isemaildisposable.webgazer.io/?email=johndoe@getnada.com` will return with this response:

```json
{
  "isDisposable": true
}
```

You can send a full e-mail address or only the domain, too. If you don't set the `email` query parameter, you will get an `HTTP 400 Bad Request` response.

## Deploying

The application is also available to be used as a [docker image](https://hub.docker.com/repository/docker/th0th/is-email-disposable). You can deploy your own self-hosted version:

```shell script
$ docker run -p 80:80 th0th/is-email-disposable
```

## I found a new domain...

If you came across a new domain that is used a disposable e-mail address source, you can add it to `domains.json` in a pull request. Please mind that domains in this file are listed alphabetically. 

## Contribution

Even if this is a very simple application, if you have some kind of improvement idea, don't hesitate to create an issue. Or a PR would be even better :)

## Shameless plug

I am an indie hacker, and I am running two services that might be useful for your business. Check them out :)

### WebGazer

[<img alt="WebGazer" src="https://user-images.githubusercontent.com/698079/162474223-f7e819c4-4421-4715-b8a2-819583550036.png" width="256" />](https://www.webgazer.io/?utm_source=github&utm_campaign=is-email-disposable-readme)

[WebGazer](https://www.webgazer.io/?utm_source=github&utm_campaign=is-email-disposable-readme) is a monitoring service that checks your website, cron jobs, or scheduled tasks on a regular basis. It notifies
you with instant alerts in case of a problem. That way, you have peace of mind about the status of your service without
manually checking it.

### PoeticMetric

[<img alt="PoeticMetric" src="https://user-images.githubusercontent.com/698079/162474946-7c4565ba-5097-4a42-8821-d087e6f56a5d.png" width="256" />](https://www.poeticmetric.com/?utm_source=github&utm_campaign=is-email-disposable-readme)

[PoeticMetric](https://www.poeticmetric.com/?utm_source=github&utm_campaign=is-email-disposable-readme) is a privacy-first, regulation-compliant, blazingly fast analytics tool.

No cookies or personal data collection. So you don't have to worry about cookie banners or GDPR, CCPA, and PECR compliance.

## License

Copyright © 2020, Gökhan Sarı. Released under the [MIT License](LICENSE).
