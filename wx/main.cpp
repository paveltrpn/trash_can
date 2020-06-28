
// Зависимости:
// $sudo apt-get install libgtk-3-dev && sudo apt-get install mesa-common-dev && sudo apt install libglu*-mesa-dev freeglut3-dev
// $sudo apt-get install libglu1-mesa-dev
//
// Сборка:
// $git clone git://github.com/wxWidgets/wxWidgets
// $cd wxWidgets
// $mkdir wx_build
// $cd wx_build
// $../configure --with-opengl --with-gtk=3
// $make -j3                    # use 3 cores. Set to the number of cores your have. 'make' uses 1 core
// $sudo make install           # some platforms require to use 'su' instead of 'sudo'
// $sudo ldconfig               # not required in each system
//
// Либо можно установить из репозиториев:
// apt-get install libwxbase3.1* \
//                 libwxgtk3.1* \
//                 wx3.1-headers \
//                 wx-common* \
//
// wxWidgets "Hello World" Program
// For compilers that support precompilation, includes "wx/wx.h"

#include <wx/wxprec.h>
#ifndef WX_PRECOMP
    #include <wx/wx.h>
#endif

class MyApp : public wxApp {
public:
    virtual bool OnInit();
};

class MyFrame : public wxFrame {
public:
    MyFrame();
private:
    void OnHello(wxCommandEvent& event);
    void OnExit(wxCommandEvent& event);
    void OnAbout(wxCommandEvent& event);
};

enum {
    ID_Hello = 1
};

wxIMPLEMENT_APP(MyApp);

bool MyApp::OnInit() {
    MyFrame *frame = new MyFrame();
    frame->Show(true);
    return true;
}

MyFrame::MyFrame() : wxFrame(NULL, wxID_ANY, "Hello World") {
    wxMenu *menuFile = new wxMenu;
    menuFile->Append(ID_Hello, "&Hello...\tCtrl-H",
                     "Help string shown in status bar for this menu item");
    menuFile->AppendSeparator();
    menuFile->Append(wxID_EXIT);
    wxMenu *menuHelp = new wxMenu;
    menuHelp->Append(wxID_ABOUT);
    wxMenuBar *menuBar = new wxMenuBar;
    menuBar->Append(menuFile, "&File");
    menuBar->Append(menuHelp, "&Help");
    SetMenuBar( menuBar );
    CreateStatusBar();
    SetStatusText("Welcome to wxWidgets!");
    Bind(wxEVT_MENU, &MyFrame::OnHello, this, ID_Hello);
    Bind(wxEVT_MENU, &MyFrame::OnAbout, this, wxID_ABOUT);
    Bind(wxEVT_MENU, &MyFrame::OnExit, this, wxID_EXIT);
}

void MyFrame::OnExit(wxCommandEvent& event) {
    Close(true);
}

void MyFrame::OnAbout(wxCommandEvent& event) {
    wxMessageBox("This is a wxWidgets Hello World example",
                 "About Hello World", wxOK | wxICON_INFORMATION);
}

void MyFrame::OnHello(wxCommandEvent& event) {
    wxLogMessage("Hello world from wxWidgets!");
}
