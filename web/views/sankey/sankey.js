'use strict';

angular.module('app.sankey', ['ngRoute', 'ngSankey'])

    .config(['$routeProvider', function($routeProvider) {
        $routeProvider.when('/sankey', {
            templateUrl: 'views/sankey/sankey.html',
            controller: 'SankeyCtrl'
        });
    }])

    .controller('SankeyCtrl', function($scope, $http, $routeParams) {
        $scope.selectedAddress = "";
        $scope.options = {
            chart: '#sankeyChart',
            width: 960,
            height: 500,
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

        $http.get(urlapi + 'alladdresses')
            .then(function(data, status, headers, config) {
                console.log('data success');
                console.log(data);
                $scope.addresses = data.data;
            }, function(data, status, headers, config) {
                console.log('data error');
            });
        $scope.getAddressSankey = function(address) {
            console.log(address);
            $scope.selectedAddress = address.id;
            $scope.data.nodes = [];
            $scope.data.links = [];
            $http.get(urlapi + 'address/sankey/' + address.id)
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
        };
    });
