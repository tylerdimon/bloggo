Convert Markdown files in `posts/` directory to HTML files in `html/` directory.
File names will be kept the same, only extensions will be changed. `first-post.md` -> `first-post.html`.

To write a file in HTML just don't create a Markdown version.
This is common practice on some pages like the home page `index.html`.

```
go build ./cmd/convert
./convert
```

Once you've generated the HTML run the server like so. Default port is `8000`

```
go build ./cmd/server
./server
```

Check out blog at `http://localhost:8080`