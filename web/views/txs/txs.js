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
        $scope.page = 1;
        $scope.count = 10;
        $scope.getTxs = function() {
            $http.get(urlapi + 'txs/' + $scope.page + '/' + $scope.count)
                .then(function(data, status, headers, config) {
                    console.log(data);
                    $scope.txs = data.data;
                }, function(data, status, headers, config) {
                    console.log('data error');
                });
        };
        $scope.getTxs();

        $scope.getPrev = function(){
                $scope.page++;
                $scope.getTxs();
        };
        $scope.getNext = function(){
                $scope.page--;
                $scope.getTxs();
        };
    });
