var tweets = new Vue({
  el: '#app9',
  data: {
    tweets: []
  },
  //mounted() {}
});

// get all tweets
var Get_tweets = new Vue({
    el: '#get_tweets',
    data: {},
    methods: {
       get_tweets: function() {
            axios.get('http://localhost:9000/tweets')
            .then(response => (tweets.tweets = response.data))
       }
    }  
})

// post new tweet
var Post_tweets = new Vue({
    el: '#post_tweets',
    data:{
        nil : false,
        tweet: {title:"new topic", }
    },
    methods: {
        post_tweets: function() { // undefined
            //console.log(this.tweet.body == undefined )
            if (this.tweet.body != undefined) {
                this.nil = false

                // hide wornning message
                // this.nil = false,

                // post tweeti
                axios.post('http://localhost:9000/tweet/new',
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

var emptyt = new Vue({
    //el : '#nil',
    methods:{

    }
})
