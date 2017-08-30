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
        $http.get(urlapi + 'lasttx')
            .then(function(data, status, headers, config) {
                console.log(data);
                $scope.txs = data.data;
            }, function(data, status, headers, config) {
                console.log('data error');
            });

    });
