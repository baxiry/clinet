var tweets = new Vue({
  el: '#app9',
  data: {
    tweets: []
  },
  mounted() {}
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
    //data:{},
    methods: {
        post_tweets: function() {
            axios.post('http://localhost:9000/tweet/new',
            {title: "new title", body: "new body"})
            .then(function (response) {
                console.log(response.data)
            })
            .catch(function (error) {
                console.log("Have an: "+ error)
            });
        }
    }
})
