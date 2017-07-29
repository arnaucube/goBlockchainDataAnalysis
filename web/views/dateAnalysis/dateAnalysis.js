'use strict';

angular.module('app.dateAnalysis', ['ngRoute', 'chart.js'])

    .config(['$routeProvider', function($routeProvider) {
        $routeProvider.when('/dateAnalysis', {
            templateUrl: 'views/dateAnalysis/dateAnalysis.html',
            controller: 'DateAnalysisCtrl'
        });
    }])

    .controller('DateAnalysisCtrl', function($scope, $http, $routeParams) {
        $scope.data=[];
        $scope.labels=[];

        $http.get(urlapi + 'houranalysis')
            .then(function(data, status, headers, config) {
                console.log('data success');
                console.log(data);

                $scope.data = data.data.data;
                $scope.labels=data.data.labels;
            }, function(data, status, headers, config) {
                console.log('data error');
            });

    });
