
var app1 = new Vue({
  el: '#app1',
  data: {
      message: 'Ok Hello Vue!'
  }
})

var app2 = new Vue({
    el: '#app2',
    data: {
       message: 'som discreptYou loaded this page on '+
        new Date().toLocaleString(),
    }
})
