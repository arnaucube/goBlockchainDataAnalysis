'use strict';

angular.module('app.dateAnalysis', ['ngRoute', 'chart.js'])

    .config(['$routeProvider', function($routeProvider) {
        $routeProvider.when('/dateAnalysis', {
            templateUrl: 'views/dateAnalysis/dateAnalysis.html',
            controller: 'DateAnalysisCtrl'
        });
    }])

    .controller('DateAnalysisCtrl', function($scope, $http) {
        $scope.last7day={
            data: [],
            labels: []
        };

        $http.get(urlapi + 'last7day')
            .then(function(data, status, headers, config) {
                console.log('data success');
                console.log(data);

                $scope.last7day.data = data.data.data;
                $scope.last7day.labels=data.data.labels;
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

        $scope.last7dayhour= {
            data:[],
            labels:  []
        };
        $http.get(urlapi + 'last7dayhour')
            .then(function(data, status, headers, config) {
                console.log('data success');
                console.log(data);

                $scope.last7dayhour.data = data.data.data;
                $scope.last7dayhour.labels = data.data.labels;
                $scope.last7dayhour.series = data.data.series;
            }, function(data, status, headers, config) {
                console.log('data error');
            });
    });
