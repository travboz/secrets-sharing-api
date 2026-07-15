# Secrets Sharing Web API

![License](https://img.shields.io/badge/license-MIT-yellow)
![go version](https://img.shields.io/badge/go-v1.26.4-blue)
<!-- ![Build Status](https://img.shields.io/badge/build-passing-brightgreen)
![Coverage](https://img.shields.io/badge/coverage-70%25-salmon)
![Downloads](https://img.shields.io/badge/downloads-0k%2Fmonth-purple) -->

> **One-liner:** Easily **create** and **share** secrets with colleagues using this small API.

<figure>
    <img src="knight-gopher.png"
         alt="Hardened Knight Gopher"
         width="350px">
    <figcaption>Our Knight Gopher standing ready to defend against any attempts at brute force attacks!.</figcaption>
</figure>

## Roadmap & Why This Project?

I purchased this project course because I was under the impression *we'd* build something cool - turns out **I** was the one doing the building. This project helped build confidence in testing and problem solving (specifically, working through the steps required to solve a problem).

See the [project](https://www.manning.com/liveproject) briefs on Manning for more information.

- [x] **Part 1, Create the Secret Sharing API:** Create a web application which will allow you to create and view secrets.
- [ ] **Part 2, Create a Custom HTTP Client:** Implement a client for the web application to enable users to create and view secrets from the command line.

### Note: My learning experience

> **Part 1:** I feel like the structure of this project may have been overengineered. The saving grace here is that I believe I can easily extend it - for example by swapping out the file-based storage with S3 or MongoDB. It is also easier tested as the interfaces allow for simple mocking. I have also attempted to implement some methods learned in Learn Go with Tests (for example, spying on calls).
> **Part 2:** Flags... what to do with flags. CLIs seem to be a common usecase for Go. There are so many packages to choose from which help streamline and expedite the creation of a CLI using the language. I jumped around a lot and was going to use urfave/cli because it seemed very intuitive, I was going to use Cobra because there was so much coverage of the tool, and everyone seems to be using it. However, as the brief wanted me to, I stuck with the `flag` package from the standard library. It was difficult, and I ended up implementing subcommands instead of picking words out of the soup of arguments like the brief said to. This is mainly because I love the `kubectl` and `git` style subcommands.

## Quick Start

```bash
# Clone repo to local machine
git clone https://github.com/travboz/secrets-sharing-api.git

# cd into repository root directory
cd secrets-sharing-api

# Pick a part then pick a milestone
cd part1/milestone2-code

# Navigate to that directory's location in the GitHub repository and read its
# 'README.md' for information on how to use it.
```

## Directory structure

To illustrate, here is the tree from the `root` directory.

The project is split into **two** parts:

1. Building the Secret Sharing Web Application API to create and view secrets.
2. Build a Secret Sharing HTTP Client CLI to create and view secrets - for a better user experience.

The other milestone directories follow a similar structure.

```bash
.
├── README.md # Info and description on the repo.
├── part1 # Web app part
│   ├── README.md # Info and description on of all relevant Part 1 code.
│   ├── knight-gopher.png # Our noble protector!
│   ├── milestone1-code # All code pertaining to the first milestone of part 1.
│   ├── milestone2-code
│   └── milestone3-code
└── part2 # CLI client part
    └── milestone1-code # All code pertaining to the first milestone of part 2.
```

## Performance

> ***NOTE:*** There is **a lot of repetition** in the `cmd/api/XXX_test.go` files. This is intended as the focus was to complete the project. Refactoring will occur later - when application is completed.

## Troubleshooting

> todo

### Still stuck?

- Open a [new issue](link)

## License

MIT © [Travis](https://github.com/travboz)

## Acknowledgments

- [Manning](https://www.manning.com/) - Producer of liveProject
- [Gopher](https://github.com/egonelbre/gophers) illustration by Egon Elbre (egonelbre/gophers), [CC0 1.0](https://creativecommons.org/publicdomain/zero/1.0/)
