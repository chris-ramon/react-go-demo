var React = require('react');
var express = require('express');
var path = require('path');
var app = express();
require('node-jsx').install();
app.get('/', function(req, res) {
  var component = require(path.resolve(req.query['module']));
  var props = JSON.parse(req.query['props'] || '{}');
  var html = React.renderToString(React.createElement(component, props));
  res.send(html);
});
app.listen(3000);
