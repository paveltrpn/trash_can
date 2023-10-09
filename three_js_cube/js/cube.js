
// Получаем указатель на элемент html в который будем рендерить
container = document.getElementById("cube");
// Получаем ширину этого элемента
var width = container.offsetWidth;
var height = 500;

// Создаём рендер
renderer = new THREE.WebGLRenderer( { antialias: true } );
renderer.setPixelRatio( window.devicePixelRatio );
renderer.setSize( width, height);

// Присоединяем внутреннюю канву с рендером к ДОМ элементу, в который
// будем рендерить
container.appendChild(renderer.domElement);

var scene = new THREE.Scene();

var cubeGeometry = new THREE.CubeGeometry(200, 200, 200);
var cubeMaterial = new THREE.MeshLambertMaterial({ color: 0xffffff });
var cube = new THREE.Mesh(cubeGeometry, cubeMaterial);
cube.rotation.y = Math.PI * 45 / 180;
cube.rotation.z = Math.PI * 45 / 180;

scene.add(cube);

var camera = new THREE.PerspectiveCamera(45, width / height, 0.1, 10000);
camera.position.y = 300;
camera.position.z = 500;
camera.lookAt(cube.position);

scene.add(camera);

var skyboxGeometry = new THREE.CubeGeometry(10000, 10000, 10000);
var skyboxMaterial = new THREE.MeshBasicMaterial({ color: 0xFFFFFFFF, side: THREE.BackSide });
var skybox = new THREE.Mesh(skyboxGeometry, skyboxMaterial);

scene.add(skybox);

var pointLight = new THREE.PointLight(0xffffff);
pointLight.position.set(0, 400, 300);

scene.add(pointLight);

var clock = new THREE.Clock();

function render() {
        requestAnimationFrame(render);
        
        cube.rotation.y -= clock.getDelta();
        cube.rotation.z -= clock.getDelta();
        
        renderer.render(scene, camera);
}

render();