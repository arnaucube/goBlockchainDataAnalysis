'use strict';

angular.module('app.main', ['ngRoute'])

    .config(['$routeProvider', function($routeProvider) {
        $routeProvider.when('/main', {
            templateUrl: 'views/main/main.html',
            controller: 'MainCtrl'
        });
    }])

    .controller('MainCtrl', function($scope, $http) {
        //last addr
        $scope.addresses = [];
        $http.get(urlapi + 'lastaddr')
            .then(function(data, status, headers, config) {
                console.log('data success');
                console.log(data);

                $scope.addresses = data.data;
            }, function(data, status, headers, config) {
                console.log('data error');
            });

        //last tx
        $scope.txs = [];
        $http.get(urlapi + 'lasttx')
            .then(function(data, status, headers, config) {
                console.log('data success');
                console.log(data);

                $scope.txs = data.data;
            }, function(data, status, headers, config) {
                console.log('data error');
            });

        //date analysis
        $scope.last24hour= {
            data:[],
            labels:  []
        };
        $http.get(urlapi + 'last24hour')
            .then(function(data, status, headers, config) {
                console.log('data success');
                console.log(data);

                $scope.last24hour.data = data.data.data;
                $scope.last24hour.labels = data.data.labels;
            }, function(data, status, headers, config) {
                console.log('data error');
            });
        $scope.last7day= {
            data:[],
            labels:  []
        };
        $http.get(urlapi + 'last7day')
            .then(function(data, status, headers, config) {
                console.log('data success');
                console.log(data);

                $scope.last7day.data = data.data.data;
                $scope.last7day.labels = data.data.labels;
            }, function(data, status, headers, config) {
                console.log('data error');
            });
    });
