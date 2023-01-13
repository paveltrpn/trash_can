
import sdl2
import sdl2.ext
import vulkan as vk
import sys

WND_WIDTH = 800
WND_HEIGHT = 600
WND_NAME = "basic"

def debugCallback(*args):
    print('DEBUG: ' + args[5] + ' ' + args[6])
    return 0

class vkAppState_c:
    #instExtNamesList: list[str]
    #instLayersNamesList: list[str]

    def __init__(self, enableLayers: bool = True) -> None:
        self._enableLayers = enableLayers
    
    def __del__(self) -> None:
        vk.vkDestroyInstance(self.instance, None)

    def getInstanceExtensionsList(self) -> None:
        extensions = vk.vkEnumerateInstanceExtensionProperties(None)
        self.instExtNamesList = [e.extensionName for e in extensions]


    def getInstanceLayersList(self) -> None:
        layers = vk.vkEnumerateInstanceLayerProperties()
        self.instLayersNamesList = [l.layerName for l in layers]

    
    def createInstance(self) -> None:
        appInfo = vk.VkApplicationInfo(
            sType=vk.VK_STRUCTURE_TYPE_APPLICATION_INFO,
            pApplicationName="basic",
            applicationVersion=vk.VK_MAKE_VERSION(1, 0, 0),
            pEngineName="No Engine",
            engineVersion=vk.VK_MAKE_VERSION(1, 0, 0),
            apiVersion=vk.VK_API_VERSION_1_0)
        
        extCount = 0
        extNames = ""
        layersCount = 0
        layersNames = ""

        if self._enableLayers == True:
            extCount = len(self.instExtNamesList)
            extNames = self.instExtNamesList
            layersCount = len(self.instLayersNamesList)
            layersNames = self.instLayersNamesList

        createInfo = vk.VkInstanceCreateInfo(
            sType=vk.VK_STRUCTURE_TYPE_INSTANCE_CREATE_INFO,
            flags=0,
            pApplicationInfo=appInfo,
            enabledExtensionCount=extCount,
            ppEnabledExtensionNames=extNames,
            enabledLayerCount=layersCount,
            ppEnabledLayerNames=layersNames)
        
        try:
            self.instance = vk.vkCreateInstance(createInfo, None)
        except:
            print("can't create vulkan instance! - ", sys.exc_info()[0])
            sys.exit(0)

        vkCreateDebugReportCallbackEXT = vk.vkGetInstanceProcAddr(
            self.instance,
            "vkCreateDebugReportCallbackEXT")
        vkDestroyDebugReportCallbackEXT = vk.vkGetInstanceProcAddr(
            self.instance,
            "vkDestroyDebugReportCallbackEXT")
        
        debug_create = vk.VkDebugReportCallbackCreateInfoEXT(
            sType=vk.VK_STRUCTURE_TYPE_DEBUG_REPORT_CALLBACK_CREATE_INFO_EXT,
            flags=vk.VK_DEBUG_REPORT_ERROR_BIT_EXT | vk.VK_DEBUG_REPORT_WARNING_BIT_EXT,
            pfnCallback=debugCallback)
        callback = vkCreateDebugReportCallbackEXT(self.instance, debug_create, None)


def main() -> None:
    sdl2.ext.init()
    window = sdl2.ext.Window(WND_NAME, size=(WND_WIDTH, WND_HEIGHT))
    window.show()

    vkState = vkAppState_c()
    
    vkState.getInstanceExtensionsList()
    print(vkState.instExtNamesList)

    vkState.getInstanceLayersList()
    print(vkState.instLayersNamesList)

    vkState.createInstance()

    run = True
    while run:
        events = sdl2.ext.get_events()
        for event in events:
            if event.type == sdl2.SDL_QUIT:
                run = False
                break
            
            if event.type == sdl2.SDL_KEYDOWN:
                if event.key.keysym.sym == sdl2.SDLK_ESCAPE:
                    run = False
                if event.key.keysym.sym == sdl2.SDLK_q:
                    run = False

        window.refresh()

    sdl2.ext.quit()

if __name__ == "__main__":
    main()
