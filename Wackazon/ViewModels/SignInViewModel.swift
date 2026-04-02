//
//  SignInViewModel.swift
//  Wackazon
//
//  Created by Oscar Leigh on 02/04/2026.
//

import SwiftUI

@Observable
class SignInViewModel {
    var usernameInput: String = ""
    var passwordInput: String = ""
    
    var navigateToHomePage: Bool = false
    
    var errorMessage: String = ""
    var showError: Bool = false
    
    func isUsernameValid() -> Bool {
        //db implementation required
        return true
    }
    
    func isPasswordValid() -> Bool {
        //db implementation required
        return true
    }
    
    func goToHomePage() {
        if (isPasswordValid() && isUsernameValid()) {
            navigateToHomePage = true
        } else {
            showError = true
        }
    }
}
