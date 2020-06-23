// Define a new component called todo-item
Vue.component('todo-item', {
    props: ['todo'],
    template: '<li>{{todo.text}}</li>',
})

var app8 = new Vue({
    el: '#app8',
    data: {
        todos:[
            {id: 0, text: "buil big progict"},
            {id: 1, text: "buil awsome progict"},
            {id: 2, text: "buil great progict"},
        ]
    }
})

//Now you can compose it in another component’s template:


var app7 = new Vue({
    el: '#app7',
    data: {},
    methods: {
        toggle: function() {
            s = app3.show
            if ( app3.show == false) {
                app3.show = true
            } else {
                app3.show = false
            } 
        }
    }

})
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
            {text: "build the big and owsome thing"}
        ]
    }
})


var app3 = new Vue({
    el: '#app3',
    data: {
        show: true,
    },
})


