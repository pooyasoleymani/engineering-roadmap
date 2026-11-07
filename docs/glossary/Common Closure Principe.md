---
Date: 2026-11-07
tags:
  - architectre
  - design_pattern
---
---
The Common Closure Principle (CCP) is a software design principle introduced by Robert C. Martin (Uncle Bob). It is one of the three principles of component cohesion, alongside:

1. Reuse/Release Equivalence Principle (REP)
2. Common Closure Principle (CCP)
3. Common Reuse Principle (CRP)

## 1. Definition

The Common Closure Principle states:

> The classes in a component should be closed together against the same kinds of changes. A change that affects a component should affect that component only.

In simpler terms:

Classes that tend to change for the same reason should be grouped together in the same component (module or package).

The goal is to minimize the number of components that must be modified when a requirement changes.

## 2. Practical example

Imagine you are developing a banking application with these classes:

```cpp
class Account {
public:
    void deposit(double amount);
    void withdraw(double amount);
};

class AccountReport {
public:
    void generateReport(const Account& account);
};

class AccountRepository {
public:
    void save(const Account& account);
};
```

Suppose the business rules for calculating account balances change.

The following classes might need modification:

- `Account`
- `AccountReport`

But `AccountRepository` probably does not need to change.

If you group classes based on their reasons for changing, you can organize them into separate components:

```
Banking/
│
├── AccountManagement/
│   └── Account
│
├── Reporting/
│   └── AccountReport
│
└── Persistence/
    └── AccountRepository
```

This is a simplified example. The actual component boundaries depend on which classes change together.

## 3. Why is it important?

Without CCP, a single business requirement might force you to modify several unrelated components.

For example:

Without CCP

One requirement change

A change to account calculation rules

Account

Modified

Reporting

Modified

Persistence

Modified

With CCP

One requirement change

Only the component responsible for account calculation changes.

Benefits:

- Reduces the impact of changes.
- Makes maintenance easier.
- Reduces unnecessary recompilation and redeployment.
- Improves component cohesion.
- Helps isolate responsibilities.

## 4. Relationship with Single Responsibility Principle (SRP)

CCP is closely related to SRP, but they operate at different levels.

| Principle | Scope     | Main idea                                                            |
| --------- | --------- | -------------------------------------------------------------------- |
| SRP       | Class     | A class should have one reason to change.                            |
| CCP       | Component | Classes should be grouped when they have the same reasons to change. |

Think of it this way:

- SRP: Which responsibilities belong inside a class?
- CCP: Which classes belong inside the same component?

## 5. Real-world example in C++

Suppose you are developing your RF Spectrum Analyzer.

You have these classes:

```
SpectrumAnalyzer/
├── FFTProcessor
├── WaterfallRenderer
├── FrequencyController
├── SCPIClient
├── DFProcessor
└── DFCompassWidget
```

Consider two types of changes:

- A change in FFT processing algorithms.
- A change in the SCPI communication protocol.

These changes have different reasons and should generally affect different components.

A possible organization:

```
SpectrumAnalyzer/
│
├── SignalProcessing/
│   ├── FFTProcessor
│   └── DFProcessor
│
├── Visualization/
│   ├── WaterfallRenderer
│   └── DFCompassWidget
│
└── Communication/
    ├── SCPIClient
    └── FrequencyController
```

This organization makes it easier to modify signal-processing algorithms without unnecessarily changing the communication or visualization components.

## 6. Test your understanding

Consider this scenario:

You have three classes:

- `User`
- `UserValidator`
- `EmailSender`

A new requirement changes the user password validation rules.

According to CCP, which classes should ideally be grouped together?

A. User and EmailSender
B. User and UserValidator
C. All three classes

User and UserValidator are likely to change together when password validation requirements change.

Key takeaway: CCP is about organizing software components around the reasons their classes change, so that related changes remain localized.