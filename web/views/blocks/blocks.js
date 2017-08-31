'use strict';

angular.module('app.blocks', ['ngRoute'])

    .config(['$routeProvider', function($routeProvider) {
        $routeProvider.when('/blocks', {
            templateUrl: 'views/blocks/blocks.html',
            controller: 'BlocksCtrl'
        });
    }])

    .controller('BlocksCtrl', function($scope, $http) {

        //last tx
        $scope.txs = [];
        $scope.page = 1;
        $scope.count = 10;
        $scope.getBlocks = function() {;
            $http.get(urlapi + 'blocks/' + $scope.page + '/' + $scope.count)
                .then(function(data, status, headers, config) {
                    console.log(data);
                    $scope.txs = data.data;
                }, function(data, status, headers, config) {
                    console.log('data error');
                });
        };
        $scope.getBlocks();

        $scope.getPrev = function(){
                $scope.page++;
                $scope.getBlocks();
        };
        $scope.getNext = function(){
                $scope.page--;
                $scope.getBlocks();
        };

    });
