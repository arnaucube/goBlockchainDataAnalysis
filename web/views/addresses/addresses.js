'use strict';

angular.module('app.addresses', ['ngRoute'])

    .config(['$routeProvider', function($routeProvider) {
        $routeProvider.when('/addresses', {
            templateUrl: 'views/addresses/addresses.html',
            controller: 'AddressesCtrl'
        });
    }])

    .controller('AddressesCtrl', function($scope, $http) {

        //last addr
        $scope.addresses = [];
        $scope.page = 1;
        $scope.count = 10;
        $scope.getAddresses = function() {
            $http.get(urlapi + 'addresses/' + $scope.page + '/' + $scope.count)
                .then(function(data, status, headers, config) {
                    console.log(data);
                    $scope.addresses = data.data;
                }, function(data, status, headers, config) {
                    console.log('data error');
                });
        };
        $scope.getAddresses();

        $scope.getPrev = function(){
                $scope.page++;
                $scope.getAddresses();
        };
        $scope.getNext = function(){
                $scope.page--;
                $scope.getAddresses();
        };
    });
