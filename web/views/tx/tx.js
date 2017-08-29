'use strict';

angular.module('app.tx', ['ngRoute'])

    .config(['$routeProvider', function($routeProvider) {
        $routeProvider.when('/tx/:txid', {
            templateUrl: 'views/tx/tx.html',
            controller: 'TxCtrl'
        });
    }])

    .controller('TxCtrl', function($scope, $http, $routeParams) {
        $scope.tx = {};
        $http.get(urlapi + 'tx/' + $routeParams.txid)
            .then(function(data, status, headers, config) {
                console.log(data);
                $scope.tx = data.data;
            }, function(data, status, headers, config) {
                console.log('data error');
            });
    });
