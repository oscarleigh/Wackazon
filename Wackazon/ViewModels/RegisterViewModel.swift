//
//  RegisterViewModel.swift
//  Wackazon
//
//  Created by Oscar Leigh on 02/04/2026.
//

import SwiftUI

@Observable
class RegisterViewModel {
    var usernameInput: String = ""
    var passwordInput: String = ""
    
    var navigateToHomePage: Bool = false
    
    var errorMessage: String = ""
    var showError: Bool = false
    
    func isUsernameValid() -> Bool {
        if (usernameInput.isEmpty) {
            errorMessage = "Please enter a username"
            return false
        }
        return true
    }
    
    func isPasswordValid() -> Bool {
        if (passwordInput.isEmpty) {
            errorMessage = "Please enter a password"
            return false
        }
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
