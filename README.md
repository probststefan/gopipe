# gopipe

I had several large datasets which had to be processed in steps. At each editing step I wanted to be free in the choice of programming language and the type of implementation.

The result is a simple Go program that starts a separate Docker container for each step and executes a command or script there.

## Pipeline example

Create a `pipeline.yml` and get started:

``` yaml
step1:
  image: node:8
  script:
    - echo
    - Hello!
step2:
  image: php:cli
  script:
    - php
    - -v
```
