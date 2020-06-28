var tweets = new Vue({
  el: '#app9',
  data: {
    tweets: []
  },
  mounted() {
      //axios.get('http://localhost:9000/tweets')
      //.then(response => (this.tweets = response.data))
  }
});

//Now you can compose it in another component’s template:
var i = 0;
var Get_tweets = new Vue({
    el: '#get_tweets',
    data: {},
    methods: {
       get_tweets: function() {
            axios.get('http://localhost:9000/tweets')
            .then(response => (tweets.tweets = response.data))
            axios.post('localhost:9000/tweet/new/idea/"ok everything is good"');
            console.log("is ok")
       }
    },
      
})

var Post_tweets = new Vue({
    el: '#post_tweets',
    data:{
        tweets:[
            {title: "new title", body: "new body"},
        ]
    },
    methods: {
        post_tweets: function() {
            axios.post('http://localhost:9000/tweet/new/ok/"some data"', {
                title: this.title,
                body: this.body,
            })
            .then(function (response) {
                currentObj.output = response.data;
            })
            .catch(function (error) {
                currentObj.output = error;
            });
        }
    }
})
