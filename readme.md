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

You can send full e-mail address or only the domain, too. If you don't set the `email` query parameter, you will get a `HTTP 400 Bad Request` response.

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

I am an indie hacker and I am running an uptime monitoring and analytics platform called [WebGazer](https://www.webgazer.io). You might want to check it out if you are running an online business and want to notice the incidents before your customers.

## License

Copyright © 2020, Gökhan Sarı. Released under the [MIT License](LICENSE).