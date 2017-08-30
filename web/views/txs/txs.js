'use strict';

angular.module('app.txs', ['ngRoute'])

    .config(['$routeProvider', function($routeProvider) {
        $routeProvider.when('/txs', {
            templateUrl: 'views/txs/txs.html',
            controller: 'TxsCtrl'
        });
    }])

    .controller('TxsCtrl', function($scope, $http) {

        //last tx
        $scope.txs = [];
        $http.get(urlapi + 'lasttx')
            .then(function(data, status, headers, config) {
                console.log(data);
                $scope.txs = data.data;
            }, function(data, status, headers, config) {
                console.log('data error');
            });

    });
