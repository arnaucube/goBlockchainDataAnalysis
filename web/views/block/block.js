'use strict';

angular.module('app.block', ['ngRoute', 'ngSankey'])

    .config(['$routeProvider', function($routeProvider) {
        $routeProvider.when('/block/:height', {
            templateUrl: 'views/block/block.html',
            controller: 'BlockCtrl'
        });
    }])

    .controller('BlockCtrl', function($scope, $http, $routeParams) {
        $scope.block = {};
        $http.get(urlapi + 'block/' + $routeParams.height)
            .then(function(data, status, headers, config) {
                console.log(data);
                $scope.block = data.data;
            }, function(data, status, headers, config) {
                console.log('data error');
            });


        //Sankey
        $scope.options = {
            chart: '#sankeyChart',
            width: 700,
            height: 280,
            margin: {top: 1, right: 1, bottom: 6, left: 1},
            node: {width: 15, padding :10, showValue: false},
            value: {format: ',.0f', unit : ''},
            dynamicLinkColor: true,
            trafficInLinks: true
        };
        $scope.data={
            nodes: [],
            links: []
        };
        $http.get(urlapi + 'block/' + $routeParams.height + '/sankey')
            .then(function(data, status, headers, config) {
                console.log('data success');
                console.log(data);
                $scope.data.nodes = data.data.nodes;
                $scope.data.links = data.data.links;
                console.log($scope.data);
                d3.selectAll("svg > *").remove();
                let chart = new d3.sankeyChart(data.data, $scope.options);
                //$scope.data = data.data;
            }, function(data, status, headers, config) {
                console.log('data error');
            });
    });
