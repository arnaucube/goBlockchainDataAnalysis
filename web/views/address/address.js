'use strict';

angular.module('app.address', ['ngRoute'])

    .config(['$routeProvider', function($routeProvider) {
        $routeProvider.when('/address/:hash', {
            templateUrl: 'views/address/address.html',
            controller: 'AddressCtrl'
        });
    }])

    .controller('AddressCtrl', function($scope, $http, $routeParams) {
        $scope.address = {};
        $http.get(urlapi + 'address/' + $routeParams.hash)
            .then(function(data, status, headers, config) {
                console.log(data);
                $scope.address = data.data;
            }, function(data, status, headers, config) {
                console.log('data error');
            });
            $scope.addresstimechart= {
                data:[],
                labels:  []
            };
            $http.get(urlapi + 'addresstimechart/' + $routeParams.hash)
                .then(function(data, status, headers, config) {
                    console.log(data);
                    $scope.addresstimechart.data = data.data.data;
                    $scope.addresstimechart.labels = data.data.labels;
                }, function(data, status, headers, config) {
                    console.log('data error');
                });
    });
