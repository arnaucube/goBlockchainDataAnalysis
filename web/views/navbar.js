'use strict';

angular.module('app.navbar', ['ngRoute'])

    .config(['$routeProvider', function($routeProvider) {
        $routeProvider.when('/navbar', {
            templateUrl: 'views/navbar/navbar.html',
            controller: 'NavbarCtrl'
        });
    }])

    .controller('NavbarCtrl', function($scope, $http, $routeParams, $location) {
        $scope.locationHash = $location.path();
        $scope.goBack = function() {
            console.log("goBack");
            window.history.back();
        };
    });
