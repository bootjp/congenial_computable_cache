let __REQUEST_URI__;
fetch(__REQUEST_URI__)
  .then(function (response) {
    if (response.status > 204) {
      return (new Error("error response"))
    }
    return response.json();
  })
  .then(function (response) {
    return (new Error("error response"))
  });