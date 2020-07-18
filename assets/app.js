// TODO make all flles one file

// definde router
//const Foo = { template: '<div>foo</div>' }
const Bar = { template: '<div>this is big baaar</div>'   }

const router = new VueRouter({
    routes: [
        { path: '/foo', component: { template: '<div>this is big foo</div>' }  },
        { path: '/bar', component: Bar }
    ] // short for `routes: routes`
})

const app = new Vue({
  router
}).$mount('#app')


var login = new Vue({
    el:'#login',
    data: {},
    methods: {
        login: function() {

        }
    }
})

var tweets = new Vue({
  el: '#tweets',
  data: {
      tweets: []
  },
  //mounted() {}
});

// get all tweets
var Get_tweets = new Vue({
    el: '#get_tweets',
    
    methods: {
       get_tweets: function() {
            axios.get('http://localhost:8000/tweets')
            .then(response => (tweets.tweets = response.data))
       }
    }  
})

// post new tweet
var Post_tweets = new Vue({
    el: '#post_tweet',
    data:{
        nil : false,
        tweet: {title:"new topic", }
    },
    methods: {
        post_tweet: function() { // undefined
            //console.log(this.tweet.body == undefined )
            if (this.tweet.body != undefined) {
                this.nil = false

                // hide wornning message
                // this.nil = false,

                // post tweeti
                axios.post('http://localhost:8000/tweet/new',
                this.tweet,
                // or post : {title: "new title", body: "new body"}
                )
                .then(function (response) {
                    //console.log(response.data)
                })
                .catch(function (error) {
                     console.log("Have an: "+ error)
                });
            } else {
                // show warning message
                this.nil = true
            }
        },
    }
})

// delete tweet
var delete_tweet = new Vue({
    el : '#del_tweet',
    data: {
        id: 20,
    },
    methods:{
        delete_tweet: function() { // undefined
            // delete tweeti
            axios.delete('http://localhost:8000/tweet/'+this.id)
            //.then(function (response) {
                //console.log(response.data)
            //})
            .catch(function (error) {
                 console.log("Have an: "+ error)
            });
        } 
    }
})


