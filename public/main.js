
$(function() {
  'use strict';

  var server = 'http://localhost:9000';

  var draw = function(data) {
    var ctx = document.getElementById('sensors').getContext('2d');
    var options = {
      animation: false,
      pointDot : false,

    };

    var lineChart = new Chart(ctx).Line(prepareData(data), options);
  };

  var prepareData = function(backendData) {
    if (!backendData) {
      console.error('No backend data given to prepareData()');
      return;
    }

    var values = backendData.values;
    var labels = [];
    var data = [];
    var i = 0;

    var step = parseInt(backendData.values.length/100);
    alert(step);

    for (var key in backendData.values) {
      var value = values[key];
      if (i%step == 0) {
        labels.push(value.t);
      }
      data.push(value.v);
      i++;
    }

    console.log(data);

    var datasets = [];

    datasets.push({
      label: "Temperature",
      data: data
    });

    return {
      labels: labels,
      datasets: datasets
    }
  };

  var formatDate = function(date) {
    if (!date) {
      console.error('No date provided to formatDate');
      return '';
    }

    var month = date.getMonth()+1;
    if (month < 10) {
      month = '0' + month;
    }
    var day = date.getDate();
    if (day < 10) {
      day = '0' + day;
    }

    return date.getFullYear() + '-' + month + '-' + day + '%20' + '00:00';
  }

  var getData = function(start, end) {
    var xhr = new XMLHttpRequest();

    var params = '?start=' + formatDate(start) + '&end=' + formatDate(end);

    xhr.open('GET', server + '/api/graph/temp.json' + params, true);
    xhr.onreadystatechange = function (event) {
      if (xhr.readyState === 4) {
        if(xhr.status === 200) {
          onData(xhr.responseText);
        } else {
          onDataError(event);
        }
      }
    };

    xhr.send(null);
  };

  var onData = function(response) {
    var data = JSON.parse(response);
    draw(data);
  };

  var onDataError = function(event) {
  };

  var old = new Date();
  old.setMonth(1);

  var n = new Date();
  n.setDate(n.getDate()+1);
  getData(old, n);
});
