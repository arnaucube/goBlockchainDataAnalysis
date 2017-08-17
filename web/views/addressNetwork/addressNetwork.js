'use strict';

angular.module('app.addressNetwork', ['ngRoute'])

    .config(['$routeProvider', function($routeProvider) {
        $routeProvider.when('/addressNetwork', {
            templateUrl: 'views/addressNetwork/addressNetwork.html',
            controller: 'AddressNetworkCtrl'
        });
    }])

    .controller('AddressNetworkCtrl', function($scope, $http, $routeParams) {
        $scope.data = [];
        $scope.addresses;
        $scope.nodes = [];
        $scope.edges = [];
        $scope.selectedNode = {};
        var nodes, edges, container, network;
        var options = {
            layout: {
                improvedLayout: false
            },
            interaction: {
                hover: true
            },
            physics: {
                stabilization: false,
                //enabled: false
            }
        };


        $scope.showMap = function() {
            var nodes = $scope.nodes;
            var edges = $scope.edges;

            var container = document.getElementById('mynetwork');
            var data = {
                nodes: nodes,
                edges: edges
            };
            network = new vis.Network(container, data, options);
            network.on("click", function(params) {
                params.event = "[original event]";
                //$scope.selectedNode = JSON.stringify(params, null, 4);
                $scope.selectedNode = params;
                console.log($scope.selectedNode);
                console.log($scope.selectedNode.nodes);
                var options = {
                    // position: {x:positionx,y:positiony}, // this is not relevant when focusing on nodes
                    scale: 1,
                    offset: {
                        x: 0,
                        y: 0
                    },
                    animation: {
                        duration: 500,
                        easingFunction: "easeInOutQuad"
                    }
                };
                network.focus($scope.selectedNode.nodes[0], options);
                //console.log('click event, getNodeAt returns: ' + this.getNodeAt(params.pointer.DOM));
            });
        };
        $http.get(urlapi + 'alladdresses')
            .then(function(data, status, headers, config) {
                console.log('data success');
                console.log(data);
                $scope.addresses = data.data;
            }, function(data, status, headers, config) {
                console.log('data error');
            });

            $scope.selectedAddress="";
        $scope.getAddressNetwork = function(address) {
            console.log(address);
            $scope.selectedAddress=address;
            $http.get(urlapi + 'address/network/' + address.id)
                .then(function(data, status, headers, config) {
                    console.log('data success');
                    console.log(data);
                    $scope.nodes = data.data.nodes;
                    $scope.edges = data.data.edges;
                    $scope.showMap();
                }, function(data, status, headers, config) {
                    console.log('data error');
                });
        };


        $scope.focusNode = function(node) {
            var options = {
                // position: {x:positionx,y:positiony}, // this is not relevant when focusing on nodes
                scale: 1,
                offset: {
                    x: 0,
                    y: 0
                },
                animation: {
                    duration: 500,
                    easingFunction: "easeInOutQuad"
                }
            };
            network.focus(node.id, options);
        };
    });
