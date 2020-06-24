/*
const tweets = new Vue({
     el: '#app9',
  data: {
    tweets: [
      {title: "the very first post", body: "lorem ipsum some test dimpsum"},
      {title: "and then there was the second", body: "lorem ipsum some test dimsum"},
      {title: "third time's a charm", body: "lorem ipsum some test dimsum"},
      {title: "four the last time", body: "lorem ipsum some test dimsum"}
    ]
  }
});
*/
var ts = new Vue({
  el: '#app9',
  data: {
    tweets: []
  },
  mounted() {
      axios.get('http://localhost:9000/tweets')
      .then(response => (this.tweets = response.data))
  }
});

async function HeadRequest() {
  let res = await axios.head('http://localhost:9000/tweets');
  let data = await axios.get('http://localhost:9000/tweets');

  console.log(`Status: ${res.status}`)
  //console.log(`Date: ${res.headers.date}`)
    for (var i= 0; i< data.data.lenght; i++){
  console.log(data.data[i].body)
    }
}

HeadRequest();

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
    },

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


