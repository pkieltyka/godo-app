var API = API || {};

(function () {
  'use strict';

  function getCookie(name) {
    var value = "; " + document.cookie;
    var parts = value.split("; " + name + "=");
    if (parts.length == 2) return parts.pop().split(";").shift();
  }

  API.defaults = function(r) {
    var token = getCookie("jwt");
    console.log("token");

    return r
      .set("Authorization", "BEARER "+token)
  }

  API.get = function(url) {
    return this.defaults(superagent.get(url));
  }

  API.post = function(url) {
    return this.defaults(superagent.post(url));
  }

  API.put = function(url) {
    return this.defaults(superagent.put(url));
  }

  API.del = function(url) {
    return this.defaults(superagent.del(url));
  }

})();
