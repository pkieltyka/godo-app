/*jshint quotmark:false */
/*jshint white:false */
/*jshint trailing:false */
/*jshint newcap:false */
var app = app || {};

(function () {
  'use strict';

  var Utils = app.Utils;
  // Generic "model" object. You can use whatever
  // framework you want. For this application it
  // may not even be worth separating this logic
  // out, but we do this to demonstrate one way to
  // separate out parts of your application.
  app.TodoModel = function (key) {
    this.key = key;
    this.todos = Utils.store(key);
    this.onChanges = [];

    // Get initial list of todos
    var that = this;
    API.get("/todos").end(function(err, res) {
    	if (err) throw err;
    	if (res.statusType != 2) return;
    	that.todos = that.todos.concat(res.body);
    	that.inform();
    })
  };

  app.TodoModel.prototype.subscribe = function (onChange) {
    this.onChanges.push(onChange);
  };

  app.TodoModel.prototype.inform = function () {
    Utils.store(this.key, this.todos);
    this.onChanges.forEach(function (cb) { cb(); });
  };

  app.TodoModel.prototype.addTodo = function (title) {
    var that = this;
    API.post("/todos")
      .send({title: title, completed: false})
      .end(function(err, res) {
      	if (err) throw err;
      	that.todos = that.todos.concat(res.body);
      	that.inform();
      })
  };

  app.TodoModel.prototype.toggleAll = function (checked) {
  	console.log("toggleAll..", checked);

    // Note: it's usually better to use immutable data structures since they're
    // easier to reason about and React works very well with them. That's why
    // we use map() and filter() everywhere instead of mutating the array or
    // todo items themselves.
    this.todos = this.todos.map(function (todo) {
      return Utils.extend({}, todo, {completed: checked});
    });

    this.inform();
  };

  app.TodoModel.prototype.toggle = function (todoToToggle) {
  	// console.log("toggle...", todoToToggle);
    this.todos = this.todos.map(function (todo) {
      return todo !== todoToToggle ?
        todo :
        Utils.extend({}, todo, {completed: !todo.completed});
    });

    this.inform();
  };

  app.TodoModel.prototype.destroy = function (todo) {
  	var that = this;
  	API.del("/todos/"+todo.id).end(function(err,res) {
  		if (err) throw err;
  		if (res.statusType != 2) return;
  		that.todos = that.todos.filter(function (candidate) {
	      return candidate !== todo;
	    });
	    that.inform();
  	})
  };

  app.TodoModel.prototype.save = function (todoToSave, text) {
    this.todos = this.todos.map(function (todo) {
      return todo !== todoToSave ? todo : Utils.extend({}, todo, {title: text});
    });
    this.inform();

  	var req = API.put("/todos/"+todoToSave.id).send({title: text});	  
  	// TODO: what to do in case of an error..?
	  req.end(function(err,res) {
  		if (err) throw err;
  		if (res.statusType != 2) return;
  	});
  };

  app.TodoModel.prototype.clearCompleted = function () {
  	// console.log("clear completed...");
    this.todos = this.todos.filter(function (todo) {
      return !todo.completed;
    });

    this.inform();
  };

})();
