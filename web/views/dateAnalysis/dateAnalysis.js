'use strict';

angular.module('app.dateAnalysis', ['ngRoute', 'chart.js'])

    .config(['$routeProvider', function($routeProvider) {
        $routeProvider.when('/dateAnalysis', {
            templateUrl: 'views/dateAnalysis/dateAnalysis.html',
            controller: 'DateAnalysisCtrl'
        });
    }])

    .controller('DateAnalysisCtrl', function($scope, $http, $routeParams) {
        $scope.totalhour={
            data: [],
            labels: []
        };

        $http.get(urlapi + 'totalhouranalysis')
            .then(function(data, status, headers, config) {
                console.log('data success');
                console.log(data);

                $scope.totalhour.data = data.data.data;
                $scope.totalhour.labels=data.data.labels;
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
    });
