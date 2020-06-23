var app6 = new Vue({
    el: '#app6',
    data: {
        message: 'username',
    }
})

var app5 = new Vue({
    el: '#app5',
    data: {
        message: 'hello every one!',
    },
    methods: {
        reverseMsg: function() {
            this.message = this.message.split('').reverse().join('');
            app4.todos.push({text: app6.message})
        }
    }
})

var app4 = new Vue({
    el: '#app4',
    data: {
        todos: [
            {text: "learn javascript"},
            {text: "learn learn golang"},
            {text: "build the big and owsome thing"}
        ]
    }
})


var app3 = new Vue({
    el: '#app3',
    data: {
        toggle: true,
    }
})

var app2 = new Vue({
    el: '#app2',
    data: {
       message: 'som discreptYou loaded this page on '+
        new Date().toLocaleString()
    }
})

var app1 = new Vue({
  el: '#app1',
  data: {
      message: 'ok hello vue!'
  }
})
