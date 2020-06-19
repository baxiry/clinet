
var app1 = new Vue({
  el: '#app1',
  data: {
      message: 'ok hello vue!'
  }
})

var app2 = new Vue({
    el: '#app2',
    data: {
       message: 'som discreptYou loaded this page on '+
        new Date().toLocaleString()
    }
})

var app3 = new Vue({
    el: '#app3',
    data: {
        toggle: true,
    }
})
