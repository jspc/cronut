cronut
==

![sketch out of cronut](doc/cover.jpg)

| who       | what |
|-----------|------|
| dockerhub | https://hub.docker.com/r/jspc/cronut/   |
| circleci  | https://circleci.com/gh/jspc/cronut   |
| licence   | MIT   |


Running
--

This tool is best run via docker, mainly for the reliance on libgit2. It may be run as per:

```bash
$ docker run -v /var/run/docker.sock:/var/run/docker.sock jspc/cronut -repo "https://github.com/org/config-repo"
```

Or

```bash
$ docker run -e DOCKER_HOST="https://example.com" jspc/cronut -repo "https://github.com/org/config-repo"
```

It requires access to the docker daemon it is going to control, naturally, and a git repo containing configuration files for the jobs it manages; files for which look like:

```yaml
---
name: my-fantastic-container
container: example/hello-world
env:
    - KEY=value
args:
    - "-v"
cron: "* * * * *"
```


Licence
--

MIT License

Copyright (c) 2017 jspc

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
