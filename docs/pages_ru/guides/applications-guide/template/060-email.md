---
title: Работа с электронной почтой
sidebar: applications-guide
permalink: documentation/guides/applications-guide/template/060-email.html
layout: guide
toc: false
---

{% filesused title="Файлы, упомянутые в главе" %}
- .helm/templates/deployment.yaml
- .helm/secret-values.yaml
- TODO название файла
{% endfilesused %}

В этой главе мы настроим в нашем базовом приложении работу с почтой.

Для того чтобы использовать почту мы предлагаем лишь один вариант - использовать внешнее API. В нашем примере это [mailgun](https://www.mailgun.com/).

{% offtopic title="А почему бы просто не установить sendmail?" %}
TODO: ответить на этот непростой вопрос
{% endofftopic %}

Внутри исходного кода подключение к API и отправка сообщения может выглядеть так:

____________
____________
____________
____________

Для работы с mailgun необходимо пробросить в ключи доступа в приложение. Для этого стоит использовать [механизм секретных переменных](#######TODO). *Вопрос работы с секретными переменными рассматривался подробнее, [когда мы делали базовое приложение](020-basic.html#secret-values-yaml)*

{% snippetcut name="secret-values.yaml (расшифрованный)" url="#" %}
```yaml
app:
  ____________
  ____________
```
{% endsnippetcut %}

После того, как значения корректно прописаны и зашифрованы — мы можем пробросить соответствующие значения в Deployment.

{% snippetcut name="deployment.yaml" url="#" %}
```yaml
        - name: ____________
          value: ____________
```
{% endsnippetcut %}

TODO: надо дать отсылку на какой-то гайд, где описано, как конкретно использовать ____________. Мало же просто его установить — надо ещё как-то юзать в коде.


<div>
    <a href="070-redis.html" class="nav-btn">Далее: Подключаем redis</a>
</div>