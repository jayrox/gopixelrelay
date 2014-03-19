var app = angular.module('albumsApp', ['ngRoute'])

.config(function($routeProvider) {
  $routeProvider
  .when('/', {
    controller: 'albumController',
    templateUrl: '/views/albums.html'
  })
})

.controller('albumController', ['$scope', '$http',
  function ($scope, $http) {
    $http.get('/albums.json').success(function(data) {
      $scope.albums = data.albums;
    });
  $scope.edit = function(album) {
    event.preventDefault();
    if($scope.albums == undefined) {
      return false
    }
    var aid = album.album
    console.log(aid)
    $scope.albums[aid].ShowEdit = ! $scope.albums[aid].ShowEdit;
    console.log($scope.albums[aid].ShowEdit)
  }
  $scope.showEdit = function(album){
    event.preventDefault();
    if($scope.albums == undefined) {
      return false
    }
    var aid = album.album

    if($scope.albums[aid].ShowEdit == undefined) {
     $scope.albums[aid].ShowEdit = true;
    }
    return ! $scope.albums[aid].ShowEdit;
   }
  $scope.save = function(album) {
    console.log("save: "+album);
  }
}]);

app.directive('albumViewer', function() {
  return {
    restrict: 'A',
    scope: {
      'album': '=',
      'albumIndex': '='
    },
    link: function (scope, elem, attrs) {

    },
    templateUrl: '/views/album.html'
  };
});

app.directive('albumEditDialog', function() {
	return {
		restrict: 'E',
    scope: {
      'album': '=',
      'albumIndex': '='
    },
    link: function (scope, elem, attrs) {
      //console.log('albumEditDialog')
      //console.log(scope.album)
    },
		templateUrl: '/views/edit.html'
	};
});



/*
.directive('albumName', function() {
    return function(scope, element, attrs) {
        element[0].value = scope.$parent.album.Name;
    };
})

.directive('albumKey', function() {
    return function(scope, element, attrs) {
        element[0].value = scope.$parent.album.Privatekey;
    };
})

.directive('albumDescription', function() {
    return function(scope, element, attrs) {
        element[0].innerhtml = scope.$parent.album.Description;
    };
})

.directive('albumPrivate', function() {
    return function(scope, element, attrs) {
        element[0].checked = scope.$parent.album.Private;
    };
})

.directive('albumId', function() {
    return function(scope, element, attrs) {
        element[0].value = scope.$parent.album.Id;
    };
})

.directive('albumAttr', function() {
  return function(scope, element, attrs) {
    element[0].value = scope.$parent.album.$attrs.albumAttr;
  };
});
*/
