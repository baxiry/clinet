var tweets = new Vue({
  el: '#app9',
  data: {
    show: true,
    tweets: []
  },
  mounted() {
      axios.get('http://localhost:9000/tweets')
      .then(response => (this.tweets = response.data))
  }
});



//Now you can compose it in another component’s template:
var app7 = new Vue({
    el: '#app7',
    data: {},
    methods: {
        toggle: function() {
            s = tweets.show
            if ( tweets.show == false) {
                tweets.show = true
            } else {
                tweets.show = false
            } 
        }
    }

})


