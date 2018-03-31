# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: guruguru.proto

require 'google/protobuf'

Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message "Member" do
    optional :name, :string, 1
    optional :host_port, :string, 2
  end
  add_message "JoinRequest" do
    optional :joining_member, :message, 1, "Member"
  end
  add_message "JoinResponse" do
    optional :ok, :bool, 1
  end
  add_message "PokeRequest" do
    optional :from_member, :message, 1, "Member"
    optional :message, :string, 2
  end
  add_message "PokeResponse" do
    optional :ok, :bool, 1
  end
  add_message "SetNextRequest" do
    optional :next_member, :message, 1, "Member"
  end
  add_message "SetNextResponse" do
    optional :ok, :bool, 1
  end
end

Member = Google::Protobuf::DescriptorPool.generated_pool.lookup("Member").msgclass
JoinRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("JoinRequest").msgclass
JoinResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("JoinResponse").msgclass
PokeRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("PokeRequest").msgclass
PokeResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("PokeResponse").msgclass
SetNextRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("SetNextRequest").msgclass
SetNextResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("SetNextResponse").msgclass
