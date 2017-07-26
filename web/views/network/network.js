'use strict';

angular.module('app.network', ['ngRoute'])

    .config(['$routeProvider', function($routeProvider) {
        $routeProvider.when('/network', {
            templateUrl: 'views/network/network.html',
            controller: 'NetworkCtrl'
        });
    }])

    .controller('NetworkCtrl', function($scope, $http, $routeParams) {
        $scope.data = [];
        $scope.nodes = [];
        $scope.edges = [];
        var nodes, edges, container;
        var options = {
            layout: {
                improvedLayout: false
            }
            /*,
                physics:{
                    //stabilization: false,
                   // enabled: false
                }*/
        };


        $scope.showMap = function() {
            var nodes = $scope.nodes;
            var edges = $scope.edges;

            var container = document.getElementById('mynetwork');
            var data = {
                nodes: nodes,
                edges: edges
            };
            var network = new vis.Network(container, data, options);
        };

        $http.get(urlapi + 'map')
            .then(function(data, status, headers, config) {
                console.log('data success');
                console.log(data);

                $scope.nodes = data.data.nodes;
                $scope.edges = data.data.edges;
                $scope.showMap();
            }, function(data, status, headers, config) {
                console.log('data error');
            });

    });
