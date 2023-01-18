# Interview Task - Web Crawler

## Introduction

The team would like you to implement a web crawler in Go. This is a
classic engineering problem that they feel covers a number of software
engineering principles whilst also being interesting for you to work on!

> _A [web crawler](https://en.wikipedia.org/wiki/Web_crawler) is the portion of
> a search engine that scans web pages looking for links and then follows them._

The exercise described below tasks you with implementing the foundation of a web
crawler. The test is an opportunity to showcase your technical ability and
highlight how you approach writing and testing services in Go.

**Treat this like writing production code, and show us your best!**

## Tasks

1. Accept a URL from the command line.
1. Fetch body data of the supplied URL.
1. Parse the body data and extract URLs from the href of all anchor tags.
1. Output all found URLs.

> _Note: the use of the standard library (including `golang.org/x/...` packages)
> is strongly encouraged. 3rd-party community packages should be avoided if
> possible (apart from testing related packages)._