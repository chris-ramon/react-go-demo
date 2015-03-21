var React = require('react');
var Todo = React.createClass({
  render: function() {
    return (
      <li>{this.props.todo.content}</li>
    );
  }
});
var Todos = React.createClass({
  render: function() {
    var todos = [];
    this.props.todos.forEach(function(todo) {
      todos.push(<Todo key={todo.id} todo={todo} />);
    });
    return (<ul>{todos}</ul>);
  }
});
var App = React.createClass({
  getInitialState: function() {
    //return {
      //todos: this.props.todos
    //}
    return {
      todos: [{id: 1, done: false, content: "run 10k nike"}]
    }
  },
  render: function() {
    return (
      <Todos todos={this.state.todos} />
    );
  }
});
module.exports = App;
