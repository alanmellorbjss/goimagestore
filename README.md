# goimagestore

Image API for hack day.

Reverse engineered from "bjssacademywebsites" version. This API works with:

- Nuxt/Vue starter
- React starter

## API

The API will run on port **8090**. It can be accessed locally as

```
http://localhost:8090
```

Supported routes are:

### Upload single PNG image

```
POST {baseUrl}/IL/teams/{team}/files
```

- {team} is the team name eg atari (case sensitive)

Post body is a JSON object:

```json
{
  "file": "<base64 image data, no MIME type>",
  "fileName": "<unique ID, any format>"
}
```

The "file" field is the base64 encoded image data. Note that this MUST NOT include the MIME type "data:image/png;base64,".

> Important: Only PNG images are supported

### Fetch all images

```
GET {baseUrl}/IL/teams/{team}/
```

- {team} is the team name eg atari (case sensitive)

This will return all images as a JSON object:

```json
{
  "data": ["data:image/png;base64,<base64 data>"]
}
```

The object has one field `data`.

This is an array of strings. Each string begins with the MIME type for base 64 encoded PNG images, has a comma separator, then has the base64 data for the image itself.

This is a suitable format to provide as the src attribute of an html <img> tag and will cause the image to be displayed.

> Important: Only PNG images are supported
