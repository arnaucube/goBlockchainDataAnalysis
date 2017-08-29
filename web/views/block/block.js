'use strict';

angular.module('app.block', ['ngRoute'])

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
    });
