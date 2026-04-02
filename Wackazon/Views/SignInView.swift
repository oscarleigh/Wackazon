//
//  SignInView.swift
//  Wackazon
//
//  Created by Oscar Leigh on 02/04/2026.
//

import SwiftUI

struct SignInView: View {
    @State private var viewModel = SignInViewModel()
    
    var body: some View {
        NavigationStack {
            VStack(spacing: 20) {
                Text("Welcome Back")
                    .font(Font.largeTitle.bold())
                    
                VStack (alignment: .leading, spacing: 10) {
                    Text("Username")
                        .opacity(0.75)
                        .padding(.leading, 15)
                        
                    HStack {
                        Image(systemName: "person.fill")
                            .foregroundStyle(.secondary)
                        
                        TextField("Ayanokoji", text: $viewModel.usernameInput)
                    }
                    .padding()
                    .glassEffect()
                    .padding(.bottom, 20)
                    
                    Text("Password")
                        .opacity(0.75)
                        .padding(.leading, 15)
                        
                    HStack {
                        Image(systemName: "lock.fill")
                            .foregroundStyle(.secondary)
                        
                        SecureField("Ultimate Sigma", text: $viewModel.passwordInput)
                    }
                    .padding()
                    .glassEffect()
                    .padding(.bottom, 20)
                }
                
                Button() {
                    viewModel.goToHomePage()
                } label : {
                    Text("Sign In")
                        .frame(maxWidth: .infinity)
                }
                .buttonStyle(.glassProminent)
                .padding(.horizontal, 50)
            }
            .padding()
            
            .alert("Registration Failed", isPresented: $viewModel.showError) {
                Button("OK", role: .cancel) {}
            } message: {
                Text(viewModel.errorMessage)
            }
            
            .navigationDestination(isPresented: $viewModel.navigateToHomePage) {
                HomePageView()
            }
        }
    }
}

#Preview {
    SignInView()
}
