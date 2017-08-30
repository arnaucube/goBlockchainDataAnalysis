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
        $http.get(urlapi + 'lastaddr')
            .then(function(data, status, headers, config) {
                console.log(data);
                $scope.addresses = data.data;
            }, function(data, status, headers, config) {
                console.log('data error');
            });

    });
