/*Package enum is a part of the library to access the Companies House REST API

Copyright 2018 Foundation for Open Software Development (www.fosdev.org)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package enum

const mortgageDescriptionsYAML string = `classificationDesc:
    'charge-description' : "Charge description"
    'nature-of-charge' : "Nature of the charge"

status:
    'outstanding' : "Outstanding"
    'fully-satisfied' : "Satisfied"
    'part-satisfied' : "Partially satisfied"
    'satisfied' : "Satisfied"

assets-ceased-released:
    'property-ceased-to-belong' : "All of the property or undertaking no longer forms part of the charge"
    'part-property-release-and-ceased-to-belong' : "Part of the property or undertaking has been released and no longer forms part of the charge"
    'part-property-released' : "Part of the property or undertaking has been released from the charge"
    'part-property-ceased-to-belong' : "Part of the property or undertaking no longer forms part of the charge"
    'whole-property-released' : "All of the property or undertaking has been released from the charge"
    'multiple-filings' : "Multiple filings relating to the property or undertaking"
    'whole-property-released-and-ceased-to-belong' : "All of the property or undertaking has been released and no longer forms part of the charge"

particular-description:
    'short-particulars' : "Short particulars"
    'charged-property-description' : "Description of property charged"
    'charged-property-or-undertaking-description' : "Description of property or undertaking"
    'brief-description' : "Brief description"

particular-flags:
    'contains_fixed_charge' : 'Contains fixed charge'
    'contains_floating_charge' : 'Contains floating charge'
    'floating_charge_covers_all' : 'Floating charge covers all the property or undertaking of the company'
    'contains_negative_pledge' : 'Contains negative pledge'
    'chargor_acting_as_bare_trustee' : 'Chargor acting as a bare trustee for the property'

secured-details-description:
    'amount-secured' : "Amount secured"
    'obligations-secured' : "Obligations secured"

alterations-description:
    'alterations-order' : "Description of the alterations to order"
    'alterations-prohibitions' : "Description of the prohibitions"

filing_type:
    'create-charge' : "Registration of a charge"
    'create-charge-with-deed-with-charles-court-order' : "Registration of a charge with Charles court order to extend (MR01)"
    'create-charge-with-deed-with-court-order' : "Registration of a charge with court order to extend (MR01)"
    'create-charge-with-deed' : "Registration of a charge (MR01)"
    'create-charge-with-deed-with-charles-court-order-limited-liability-partnership' : "Registration of a charge with Charles court order to extend (LLMR01)"
    'create-charge-with-deed-with-court-order-limited-liability-partnership' : "Registration of a charge with court order to extend (LLMR01)"
    'create-charge-with-deed-limited-liability-partnership' : "Registration of a charge (LLMR01)"
    'acquire-charge-with-deed-with-charles-court-order' : "Registration of an acquisition with Charles court order to extend (MR02)"
    'acquire-charge-with-deed-with-court-order' : "Registration of an acquisition with court order to extend (MR02)"
    'acquire-charge-with-deed' : "Registration of an acquisition (MR02)"
    'acquire-charge-with-deed-with-charles-court-order-limited-liability-partnership' : "Registration of an acquisition with Charles court order to extend (LLMR02)"
    'acquire-charge-with-deed-with-court-order-limited-liability-partnership' : "Registration of an acquisition with court order to extend (LLMR02)"
    'acquire-charge-with-deed-limited-liability-partnership' : "Registration of an acquisition (LLMR02)"
    'debenture-with-deed-with-charles-court-order' : "Registration of a series of debentures with Charles court order to extend (MR03)"
    'debenture-with-deed-with-court-order' : "Registration of a series of debentures with court order to extend (MR03)"
    'debenture-with-deed' : "Registration of a series of debentures (MR03)"
    'debenture-with-deed-with-charles-court-order-limited-liability-partnership' : "Registration of a series of debentures with Charles court order to extend (LLMR03)"
    'debenture-with-deed-with-court-order-limited-liability-partnership' : "Registration of a series of debentures with court order to extend (LLMR03)"
    'debenture-with-deed-limited-liability-partnership' : "Registration of a series of debentures (LLMR03)"
    'charge-satisfaction' : "Satisfaction of a charge (MR04)"
    'charge-satisfaction-limited-liability-partnership' : "Satisfaction of a charge (LLMR04)"
    'charge-whole-release' : "All of the property or undertaking has been released from the charge (MR05)"
    'charge-whole-cease' : "All of the property or undertaking no longer forms part of the charge (MR05)"
    'charge-whole-both' : "All of the property or undertaking has been released and no longer forms part of the charge (MR05)"
    'charge-part-release' : "Part of the property or undertaking has been released from the charge (MR05)"
    'charge-part-cease' : "Part of the property or undertaking no longer forms part of the charge (MR05)"
    'charge-part-both' : "Part of the property or undertaking has been released and no longer forms part of the charge (MR05)"
    'charge-release-cease' : "Property or undertaking has been released or no longer forms part of the charge (MR05)"
    'charge-whole-release-limited-liability-partnership' : "All of the property or undertaking has been released from the charge (LLMR05)"
    'charge-whole-cease-limited-liability-partnership' : "All of the property or undertaking no longer forms part of the charge (LLMR05)"
    'charge-whole-both-limited-liability-partnership' : "All of the property or undertaking has been released and no longer forms part of the charge (LLMR05)"
    'charge-part-release-limited-liability-partnership' : "Part of the property or undertaking has been released from the charge (LLMR05)"
    'charge-part-cease-limited-liability-partnership' : "Part of the property or undertaking no longer forms part of the charge (LLMR05)"
    'charge-part-both-limited-liability-partnership' : "Part of the property or undertaking has been released and no longer forms part of the charge (LLMR05)"
    'charge-release-cease-limited-liability-partnership' : "Property or undertaking has been released or no longer forms part of the charge (LLMR05)"
    'trustee-acting' : "Statement of company acting as a trustee on a charge (MR06)"
    'trustee-acting-as-limited-liability-partnership' : "Statement of company acting as a trustee on a charge (LLMR06)"
    'alter-charge' : "Particulars of an instrument of alteration to a charge (MR07)"
    'alter-charge-limited-liability-partnership' : "Particulars of an instrument of alteration to a charge (LLMR07)"
    'create-charge-without-deed-with-charles-court-order' : "Registration of a charge with Charles court order to extend without a deed (MR08)"
    'create-charge-without-deed-with-court-order' : "Registration of a charge with court order to extend without a deed (MR08)"
    'create-charge-without-deed' : "Registration of a charge without a deed (MR08)"
    'create-charge-without-deed-with-charles-court-order-limited-liability-partnership' : "Registration of a charge with Charles court order to extend without a deed (LLMR08)"
    'create-charge-without-deed-with-court-order-limited-liability-partnership' : "Registration of a charge with court order to extend without a deed (LLMR08)"
    'create-charge-without-deed-limited-liability-partnership' : "Registration of a charge without a deed (LLMR08)"
    'acquire-charge-without-deed-with-charles-court-order' : "Registration of an acquisition with Charles court order to extend without a deed(MR09)"
    'acquire-charge-without-deed-with-court-order' : "Registration of an acquisition with court order to extend without a deed (MR09)"
    'acquire-charge-without-deed' : "Registration of an acquisition without a deed (MR09)"
    'acquire-charge-without-deed-with-charles-court-order-limited-liability-partnership' : "Registration of an acquisition with Charles court order to extend without a deed (LLMR09)"
    'acquire-charge-without-deed-with-court-order-limited-liability-partnership' : "Registration of an acquisition with court order to extend without a deed (LLMR09)"
    'acquire-charge-without-deed-limited-liability-partnership' : "Registration of an acquisition without a deed (LLMR09)"
    'debenture-without-deed-with-charles-court-order' : "Registration of a series of debentures with Charles court order without a deed (MR10)"
    'debenture-without-deed-with-court-order' : "Registration of a series of debentures with court order without a deed (MR10)"
    'debenture-without-deed' : "Registration of a series of debentures without a deed (MR10)"
    'debenture-without-deed-with-charles-court-order-limited-liability-partnership' : "Registration of a series of debentures with Charles court order without a deed (LLMR10)"
    'debenture-without-deed-with-court-order-limited-liability-partnership' : "Registration of a series of debentures with court order without a deed (LLMR10)"
    'debenture-without-deed-limited-liability-partnership' : "Registration of a series of debentures without a deed (LLMR10)"
    'alter-floating-charge' : "Alteration to a floating charge (466 Scot)"
    'alter-floating-charge-limited-liability-partnership' : "Alteration to a floating charge (LLP466 Scot)"
    'create-charge-pre-april-2013' : "Registration of a charge (MG01)"
    'create-charge-scotland-pre-april-2013' : "Registration of a charge (MG01s)"
    'create-charge-pre-april-2013-limited-liability-partnership' : "Registration of a charge (LLMG01)"
    'create-charge-pre-2006-companies-act' : "Registration of a charge (395)"
    'create-charge-pre-2006-companies-act-limited-liability-partnership' : "Registration of a charge (LLP395)"
    'create-charge-scotland-pre-2006-companies-act' : "Registration of a charge (410)"
    'charge-satisfaction-pre-april-2013' : "Statement of satisfaction of a charge in full or part (MG02)"
    'charge-satisfaction-scotland-pre-april-2013' : "Statement of satisfaction of a charge in full or part (MG02s)"
    'charge-satisfaction-pre-2006-companies-act' : "Statement of satisfaction of a charge in full or part (403a)"
    'charge-satisfaction-scotland-pre-2006-companies-act' : "Statement of satisfaction of a charge in full or part (419a)"
    'charge-satisfaction-pre-april-2013-limited-liability-partnership' : "Statement of satisfaction of a charge in full or part (LLMG02)"
    'charge-satisfaction-floating-charge-in-scotland-pre-april-2013' : "Statement of satisfaction of a floating charge (MG03s)"
    'charge-part-or-whole-release-pre-april-2013' : "Statement that part or the whole of the property charged has been released (MG04)"
    'charge-part-or-whole-release-scotland-pre-april-2013' : "Statement that part or the whole of the property charged has been released (MG04s)"
    'charge-part-or-whole-release-limited-liability-partnership-pre-april-2013' : "Statement that part or the whole of the property charged has been released (LLMG04)"
    'charge-released-floating-charge-in-scotland-pre-april-2013' : "Statement that part or whole of property from a floating charge has been released (MG05s)"
    'charge-released-floating-charge-in-pre-2006-companies-act' : "Statement that part or whole of property from a floating charge has been released (403b)"
    'charge-released-floating-charge-in-scotland-pre-2006-companies-act' : "Statement that part or whole of property from a floating charge has been released (419b)"
    'acquire-charge-pre-april-2013' : "Registration of an acquisition (MG06)"
    'acquire-charge-scotland-pre-april-2013' : "Registration of an acquisition (MG06s)"
    'acquire-charge-limited-liability-partnership-pre-april-2013' : "Registration of an acquisition (LLMG06)"
    'acquire-charge-pre-2006-companies-act' : "Registration of an acquisition (400)"
    'acquire-charge-scotland-pre-2006-companies-act' : "Registration of an acquisition (416)"
    'create-charge-series-debentures-pre-april-2013' : "Registration of a charge to secure a series of debentures (MG07)"
    'create-charge-series-debentures-scotland-pre-april-2013' : "Registration of a charge to secure a series of debentures (MG07s)"
    'create-charge-series-debentures-limited-liability-partnership' : "Registration of a charge to secure a series of debentures (LLMG07)"
    'create-charge-series-debentures-pre-2006-companies-act' : "Registration of a charge to secure a series of debentures (397)"
    'create-charge-series-debentures-scotland-pre-2006-companies-act' : "Registration of a charge to secure a series of debentures (413)"
    'create-issue-out-of-series-debentures-scotland-pre-2006-companies-act' : "Registration of an issue out of a series of debentures (413a)"
    'create-issue-of-secured-debentures-pre-april-2013' : "Registration of an issue of secured debentures in a series (MG08)"
    'create-issue-of-secured-debentures-scotland-pre-april-2013' : "Registration of an issue of secured debentures in a series (MG08s)"
    'create-issue-of-secured-debentures-pre-2006-companies-act' : "Registration of an issue of secured debentures in a series (397a)"
    'create-issue-of-secured-debentures-limited-liability-partnership-pre-april-2013' : "Registration of an issue of secured debentures in a series (LLMG08)"
    'create-issue-of-secured-debentures-limited-liability-partnership-scotland-pre-april-2013' : "Registration of an issue of secured debentures in a series (LLMG08s)"
    'liquidation-appointment-of-receiver' : "Appointment of a receiver or manager (RM01)"
    'liquidation-appointment-of-receiver-pre-april-2013' : "Appointment of a receiver or manager (LQ01)"
    'liquidation-appointment-of-receiver-limited-liability-partnership' : "Appointment of a receiver or manager (LLRM01)"
    'liquidation-appointment-of-receiver-scotland' : "Appointment of a receiver or manager (1 Scot)"
    'liquidation-appointment-of-receiver-by-court-scotland' : "Appointment of a receiver or manager by the court (2 Scot)"
    'liquidation-appointment-of-receiver-pre-2006-companies-act' : "Appointment of a receiver or manager (405 (1))"
    'liquidation-appointment-of-liquidator' : "Appointment of a liquidator (VL1)"
    'liquidation-cease-to-act-receiver' : "Notice of ceasing to act as a receiver or manager (RM02)"
    'liquidation-cease-to-act-receiver-pre-april-2013' : "Notice of ceasing to act as a receiver or manager (LQ02)"
    'liquidation-cease-to-act-receiver-limited-liability-partnership' : "Notice of ceasing to act as a receiver or manager (LLRM02)"
    'liquidation-cease-to-act-receiver-pre-2006-companies-act' : "Notice of ceasing to act as a receiver or manager (405 (2))"
    'liquidation-ceasing-to-act-receiver-scotland' : "Notice of ceasing to act as a receiver or manager (3 Scot)"
    'create-charge-northern-ireland-pre-2006-companies-act-form-402' : "Registration of a charge (402 NI)"
    'create-charge-northern-ireland-pre-2006-companies-act-form-402DF' : "Registration of a charge (402DF NI )"
    'create-charge-northern-ireland-pre-2006-companies-act-form-402PP' : "Registration of a charge (402PP NI)"
    'create-charge-northern-ireland-pre-2006-companies-act-form-402R' : "Registration of a charge (402R NI)"
    'create-charge-northern-ireland-pre-2006-companies-act-form-407' : "Registration of a charge (407 NI)"
    'create-charge-by-judgement-enforcement-office-northern-ireland-pre-2006-companies-act-form-408' : "Registration of a charge by the Judgement Enforcement office (408 NI)"
    'supporting-evidence-create-charge-northern-ireland-pre-2006-companies-act' : "Supporting evidence for a registration of a charge (405 NI)"
    'charge-satisfaction-northern-ireland-pre-2006-companies-act' : "Statement of satisfaction of a charge in full or part (411A NI)"
    'charge-release-cease-northern-ireland-pre-2006-companies-act' : "Statement of cease or release in full or part from a charge (411B NI)"
    'charge-part-satisfaction' : "Satisfaction of a charge in part (MR04)"
    'charge-part-satisfaction-limited-liability-partnership' : "Satisfaction of a charge in part (LLMR04)"
    'charge-satisfaction-cease-release-northen-pre-2006-companies-act-ireland' : "Satisfaction or cease or release of a charge (49a NI)"
    'create-floating-charge-scotland' : "Registration of a floating charge (MG10s)"
    'create-charge-pre-april-2013-limited-liability-partnership-scotland' : "Registration of a charge (LLMG01s)"
    'charge-satisfaction-pre-april-2013-limited-liability-partnership-scotland' : "Statement of satisfaction of a charge in full or part (LLMG02s)"
    'charge-satisfaction-floating-charge-in-scotland-pre-april-2013-limited-liability-partnership' : "Statement of satisfaction of a floating charge (LLMG03s)"
    'charge-part-or-whole-release-pre-april-2013--limited-liability-partnership-scotland' : "Statement that part or the whole of the property charged has been released (LLMG04s)"
    'charge-released-floating-charge-in-scotland-pre-april-2013--limited-liability-partnership' : "Statement that part or whole of property from a floating charge has been released (LLMG05s)"
    'acquire-charge-limited-liability-partnership-pre-april-2013-scotland' : "Registration of an acquisition (LLMG06s)"
    'create-charge-series-debentures-limited-liability-partnership-scotland' : "Registration of a charge to secure a series of debentures (LLMG07s)"
    'charge-satisfaction-pre-2006-companies-act-limited-liability-partnership' : "Statement of satisfaction of a charge in full or part (LLP403a)"
    'charge-released-floating-charge-in-pre-2006-companies-act-limited-liability-partnership' : "Statement that part or whole of property from a floating charge has been released (LLP403b)"
    'charge-satisfaction-scotland-pre-2006-companies-act-limited-liability-partnership' : "Statement of satisfaction of charge in full or part (LLP419a)"
    'charge-released-floating-charge-in-scotland-pre-2006-companies-act-limited-liability-partnership' : "Statement that part or whole of property from a floating charge has been released (LLP419b)"
    'liquidation-appointment-of-receiver-limited-liability-partnership' : "Appointment of receiver or manager (LLRM01)"
    'liquidation-cease-to-act-receiver-limited-liability-partnership' : "Notice of ceasing to act as a receiver or manager (LLRM02)"
    'liquidation-appointment-of-receiver-pre-april-2013-limited-liability-partnership' : "Appointment of a receiver or manager (LLLQ01)"
    'liquidation-cease-to-act-receiver-pre-april-2013-limited-liability-partnership' : "Notice of ceasing to act as a receiver or manager (LLLQ02)"
    'liquidation-appointment-of-receiver-pre-april-2013-overseas' : "Appointment of a receiver or manager (OSLQ01)"
    'liquidation-cease-to-act-receiver-pre-april-2013-overseas' : "Notice of ceasing to act as a receiver or manager (OSLQ02)"
    'supporting-evidence-create-charge-northern-ireland-pre-2006-companies-act-limited-liability-partnership' : "Supporting evidence for a registration of a charge (LLP405 NI)"
    'charge-satisfaction-northern-ireland-pre-2006-companies-act-limited-liability-partnership' : "Statement of satisfaction of a charge in full or part (LLP411A NI)"
    'charge-release-cease-northern-ireland-pre-2006-companies-act-limited-liability-partnership' : "Statement of cease or release in full or part from a charge (LLP411B NI)"
    'create-charge-northern-ireland-pre-2006-companies-act-form-402-limited-liability-partnership' : "Registration of a charge (LLP402 NI)"
    'create-charge-northern-ireland-pre-2006-companies-act-form-C402' : "Registration of a charge (C402 NI)"
    'create-charge-by-judgement-enforcement-office-northern-ireland-pre-2006-companies-act-form-C408' : "Registration of a charge by the Judgement Enforcement office (C408 NI)"
    'charge-certificate-northern-ireland-pre-2006-companies-act-form-C404' : "Certificate of a charge (C404 NI)"`
