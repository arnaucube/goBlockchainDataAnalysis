var urlapi = "http://127.0.0.1:3014/";

//var urlapi = document.location.href + "api/";
console.log(urlapi);

var app = angular.module("webApp", ['chart.js']);
var nodes, edges, container;
var options = {
  layout:{
    improvedLayout: false
  }/*,
    physics:{
        //stabilization: false,
       // enabled: false
    }*/
};
app.controller("webCtrl", function($scope, $http) {
    //chart
    $scope.labels=[];
    $scope.data=[];
    $scope.nodes=[];
    $scope.edges=[];

    $http.get(urlapi + 'map')
    .then(function (data) {
      console.log('data success');
      console.log(data); // for browser console
      $scope.nodes=data.data.Nodes;
      $scope.edges=data.data.Edges;
      console.log($scope.nodes);
      console.log($scope.edges);
      $scope.showMap();
        //alert("Ara mateix es mostren (entre persones i tweets): " + nodes.length + " nodes.");
        //$scope.refreshChart();
    }, function(data){
        console.log('data error');
        console.log(status);
        console.log(data);
    });

    $scope.showMap=function(){
      var nodes = $scope.nodes;
      var edges = $scope.edges;

      container = document.getElementById('mynetwork');
      var data = {
        nodes: nodes,
        edges: edges
      };
      var network = new vis.Network(container, data, options);
      toastr.info("map completed");
    };
});
